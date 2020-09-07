"""test OAuth Provider flow"""
from time import sleep

from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from structlog import get_logger

from docker import DockerClient, from_env
from docker.models.containers import Container
from docker.types import Healthcheck
from e2e.utils import USER, SeleniumTestCase
from passbook.core.models import Application
from passbook.flows.models import Flow
from passbook.policies.expression.models import ExpressionPolicy
from passbook.policies.models import PolicyBinding
from passbook.providers.oauth2.generators import (
    generate_client_id,
    generate_client_secret,
)
from passbook.providers.oauth2.models import ClientTypes, OAuth2Provider, ResponseTypes

LOGGER = get_logger()


class TestProviderOAuth2Github(SeleniumTestCase):
    """test OAuth Provider flow"""

    def setUp(self):
        self.client_id = generate_client_id()
        self.client_secret = generate_client_secret()
        self.container = self.setup_client()
        super().setUp()

    def setup_client(self) -> Container:
        """Setup client grafana container which we test OAuth against"""
        client: DockerClient = from_env()
        container = client.containers.run(
            image="grafana/grafana:7.1.0",
            detach=True,
            network_mode="host",
            auto_remove=True,
            healthcheck=Healthcheck(
                test=["CMD", "wget", "--spider", "http://localhost:3000"],
                interval=5 * 100 * 1000000,
                start_period=1 * 100 * 1000000,
            ),
            environment={
                "GF_AUTH_GITHUB_ENABLED": "true",
                "GF_AUTH_GITHUB_ALLOW_SIGN_UP": "true",
                "GF_AUTH_GITHUB_CLIENT_ID": self.client_id,
                "GF_AUTH_GITHUB_CLIENT_SECRET": self.client_secret,
                "GF_AUTH_GITHUB_SCOPES": "user:email,read:org",
                "GF_AUTH_GITHUB_AUTH_URL": self.url(
                    "passbook_providers_oauth2_github:github-authorize"
                ),
                "GF_AUTH_GITHUB_TOKEN_URL": self.url(
                    "passbook_providers_oauth2_github:github-access-token"
                ),
                "GF_AUTH_GITHUB_API_URL": self.url(
                    "passbook_providers_oauth2_github:github-user"
                ),
                "GF_LOG_LEVEL": "debug",
            },
        )
        while True:
            container.reload()
            status = container.attrs.get("State", {}).get("Health", {}).get("Status")
            if status == "healthy":
                return container
            LOGGER.info("Container failed healthcheck")
            sleep(1)

    def tearDown(self):
        self.container.kill()
        super().tearDown()

    def test_authorization_consent_implied(self):
        """test OAuth Provider flow (default authorization flow with implied consent)"""
        sleep(1)
        # Bootstrap all needed objects
        authorization_flow = Flow.objects.get(
            slug="default-provider-authorization-implicit-consent"
        )
        provider = OAuth2Provider.objects.create(
            name="grafana",
            client_id=self.client_id,
            client_secret=self.client_secret,
            client_type=ClientTypes.CONFIDENTIAL,
            response_type=ResponseTypes.CODE,
            redirect_uris="http://localhost:3000/login/github",
            authorization_flow=authorization_flow,
        )
        Application.objects.create(
            name="Grafana", slug="grafana", provider=provider,
        )

        self.driver.get("http://localhost:3000")
        self.driver.find_element(By.CLASS_NAME, "btn-service--github").click()
        self.driver.find_element(By.ID, "id_uid_field").click()
        self.driver.find_element(By.ID, "id_uid_field").send_keys(USER().username)
        self.driver.find_element(By.ID, "id_uid_field").send_keys(Keys.ENTER)
        self.driver.find_element(By.ID, "id_password").send_keys(USER().username)
        self.driver.find_element(By.ID, "id_password").send_keys(Keys.ENTER)

        self.wait_for_url("http://localhost:3000/?orgId=1")
        self.driver.find_element(By.XPATH, "//a[contains(@href, '/profile')]").click()
        self.assertEqual(
            self.driver.find_element(By.CLASS_NAME, "page-header__title").text,
            USER().username,
        )
        self.assertEqual(
            self.driver.find_element(By.CSS_SELECTOR, "input[name=name]").get_attribute(
                "value"
            ),
            USER().username,
        )
        self.assertEqual(
            self.driver.find_element(
                By.CSS_SELECTOR, "input[name=email]"
            ).get_attribute("value"),
            USER().email,
        )
        self.assertEqual(
            self.driver.find_element(
                By.CSS_SELECTOR, "input[name=login]"
            ).get_attribute("value"),
            USER().username,
        )

    def test_authorization_consent_explicit(self):
        """test OAuth Provider flow (default authorization flow with explicit consent)"""
        sleep(1)
        # Bootstrap all needed objects
        authorization_flow = Flow.objects.get(
            slug="default-provider-authorization-explicit-consent"
        )
        provider = OAuth2Provider.objects.create(
            name="grafana",
            client_id=self.client_id,
            client_secret=self.client_secret,
            client_type=ClientTypes.CONFIDENTIAL,
            response_type=ResponseTypes.CODE,
            redirect_uris="http://localhost:3000/login/github",
            authorization_flow=authorization_flow,
        )
        app = Application.objects.create(
            name="Grafana", slug="grafana", provider=provider,
        )

        self.driver.get("http://localhost:3000")
        self.driver.find_element(By.CLASS_NAME, "btn-service--github").click()
        self.driver.find_element(By.ID, "id_uid_field").click()
        self.driver.find_element(By.ID, "id_uid_field").send_keys(USER().username)
        self.driver.find_element(By.ID, "id_uid_field").send_keys(Keys.ENTER)
        self.driver.find_element(By.ID, "id_password").send_keys(USER().username)
        self.driver.find_element(By.ID, "id_password").send_keys(Keys.ENTER)

        self.assertIn(
            app.name,
            self.driver.find_element(
                By.XPATH, "/html/body/div[2]/div/main/div/form/div[2]/p[1]"
            ).text,
        )
        self.assertEqual(
            "GitHub Compatibility: Access you Email addresses",
            self.driver.find_element(
                By.XPATH, "/html/body/div[2]/div/main/div/form/div[2]/ul/li[1]"
            ).text,
        )
        sleep(1)
        self.driver.find_element(By.CSS_SELECTOR, "[type=submit]").click()

        self.wait_for_url("http://localhost:3000/?orgId=1")
        self.driver.find_element(By.XPATH, "//a[contains(@href, '/profile')]").click()
        self.assertEqual(
            self.driver.find_element(By.CLASS_NAME, "page-header__title").text,
            USER().username,
        )
        self.assertEqual(
            self.driver.find_element(By.CSS_SELECTOR, "input[name=name]").get_attribute(
                "value"
            ),
            USER().username,
        )
        self.assertEqual(
            self.driver.find_element(
                By.CSS_SELECTOR, "input[name=email]"
            ).get_attribute("value"),
            USER().email,
        )
        self.assertEqual(
            self.driver.find_element(
                By.CSS_SELECTOR, "input[name=login]"
            ).get_attribute("value"),
            USER().username,
        )

    def test_denied(self):
        """test OAuth Provider flow (default authorization flow, denied)"""
        sleep(1)
        # Bootstrap all needed objects
        authorization_flow = Flow.objects.get(
            slug="default-provider-authorization-explicit-consent"
        )
        provider = OAuth2Provider.objects.create(
            name="grafana",
            client_id=self.client_id,
            client_secret=self.client_secret,
            client_type=ClientTypes.CONFIDENTIAL,
            response_type=ResponseTypes.CODE,
            redirect_uris="http://localhost:3000/login/github",
            authorization_flow=authorization_flow,
        )
        app = Application.objects.create(
            name="Grafana", slug="grafana", provider=provider,
        )

        negative_policy = ExpressionPolicy.objects.create(
            name="negative-static", expression="return False"
        )
        PolicyBinding.objects.create(target=app, policy=negative_policy, order=0)

        self.driver.get("http://localhost:3000")
        self.driver.find_element(By.CLASS_NAME, "btn-service--github").click()
        self.driver.find_element(By.ID, "id_uid_field").click()
        self.driver.find_element(By.ID, "id_uid_field").send_keys(USER().username)
        self.driver.find_element(By.ID, "id_uid_field").send_keys(Keys.ENTER)
        self.driver.find_element(By.ID, "id_password").send_keys(USER().username)
        self.driver.find_element(By.ID, "id_password").send_keys(Keys.ENTER)
        self.wait_for_url(self.url("passbook_flows:denied"))
        self.assertEqual(
            self.driver.find_element(By.CSS_SELECTOR, "header > h1").text,
            "Permission denied",
        )
