"""test SAML Provider flow"""
from sys import platform
from time import sleep
from unittest.case import skipUnless

from docker import DockerClient, from_env
from docker.models.containers import Container
from docker.types import Healthcheck
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from structlog import get_logger

from e2e.utils import USER, SeleniumTestCase
from passbook.core.models import Application
from passbook.crypto.models import CertificateKeyPair
from passbook.flows.models import Flow
from passbook.policies.expression.models import ExpressionPolicy
from passbook.policies.models import PolicyBinding
from passbook.providers.saml.models import (
    SAMLBindings,
    SAMLPropertyMapping,
    SAMLProvider,
)

LOGGER = get_logger()


@skipUnless(platform.startswith("linux"), "requires local docker")
class TestProviderSAML(SeleniumTestCase):
    """test SAML Provider flow"""

    container: Container

    def setup_client(self, provider: SAMLProvider) -> Container:
        """Setup client saml-sp container which we test SAML against"""
        client: DockerClient = from_env()
        container = client.containers.run(
            image="beryju/saml-test-sp",
            detach=True,
            network_mode="host",
            auto_remove=True,
            healthcheck=Healthcheck(
                test=["CMD", "wget", "--spider", "http://localhost:9009/health"],
                interval=5 * 100 * 1000000,
                start_period=1 * 100 * 1000000,
            ),
            environment={
                "SP_ENTITY_ID": provider.issuer,
                "SP_SSO_BINDING": "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST",
                "SP_METADATA_URL": (
                    self.url(
                        "passbook_providers_saml:metadata",
                        application_slug=provider.application.slug,
                    )
                ),
            },
        )
        while True:
            container.reload()
            status = container.attrs.get("State", {}).get("Health", {}).get("Status")
            if status == "healthy":
                return container
            LOGGER.info("Container failed healthcheck")
            sleep(1)

    def test_sp_initiated_implicit(self):
        """test SAML Provider flow SP-initiated flow (implicit consent)"""
        # Bootstrap all needed objects
        authorization_flow = Flow.objects.get(
            slug="default-provider-authorization-implicit-consent"
        )
        provider: SAMLProvider = SAMLProvider.objects.create(
            name="saml-test",
            acs_url="http://localhost:9009/saml/acs",
            audience="passbook-e2e",
            issuer="passbook-e2e",
            sp_binding=SAMLBindings.POST,
            authorization_flow=authorization_flow,
            signing_kp=CertificateKeyPair.objects.first(),
        )
        provider.property_mappings.set(SAMLPropertyMapping.objects.all())
        provider.save()
        Application.objects.create(
            name="SAML", slug="passbook-saml", provider=provider,
        )
        self.container = self.setup_client(provider)
        self.driver.get("http://localhost:9009")
        self.driver.find_element(By.ID, "id_uid_field").click()
        self.driver.find_element(By.ID, "id_uid_field").send_keys(USER().username)
        self.driver.find_element(By.ID, "id_uid_field").send_keys(Keys.ENTER)
        self.driver.find_element(By.ID, "id_password").send_keys(USER().username)
        self.driver.find_element(By.ID, "id_password").send_keys(Keys.ENTER)
        self.wait_for_url("http://localhost:9009/")
        self.assertEqual(
            self.driver.find_element(By.XPATH, "/html/body/pre").text,
            f"Hello, {USER().name}!",
        )

    def test_sp_initiated_explicit(self):
        """test SAML Provider flow SP-initiated flow (explicit consent)"""
        # Bootstrap all needed objects
        authorization_flow = Flow.objects.get(
            slug="default-provider-authorization-explicit-consent"
        )
        provider: SAMLProvider = SAMLProvider.objects.create(
            name="saml-test",
            acs_url="http://localhost:9009/saml/acs",
            audience="passbook-e2e",
            issuer="passbook-e2e",
            sp_binding=SAMLBindings.POST,
            authorization_flow=authorization_flow,
            signing_kp=CertificateKeyPair.objects.first(),
        )
        provider.property_mappings.set(SAMLPropertyMapping.objects.all())
        provider.save()
        app = Application.objects.create(
            name="SAML", slug="passbook-saml", provider=provider,
        )
        self.container = self.setup_client(provider)
        self.driver.get("http://localhost:9009")
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
        sleep(1)
        self.driver.find_element(By.CSS_SELECTOR, "[type=submit]").click()
        self.wait_for_url("http://localhost:9009/")
        self.assertEqual(
            self.driver.find_element(By.XPATH, "/html/body/pre").text,
            f"Hello, {USER().name}!",
        )

    def test_idp_initiated_implicit(self):
        """test SAML Provider flow IdP-initiated flow (implicit consent)"""
        # Bootstrap all needed objects
        authorization_flow = Flow.objects.get(
            slug="default-provider-authorization-implicit-consent"
        )
        provider: SAMLProvider = SAMLProvider.objects.create(
            name="saml-test",
            acs_url="http://localhost:9009/saml/acs",
            audience="passbook-e2e",
            issuer="passbook-e2e",
            sp_binding=SAMLBindings.POST,
            authorization_flow=authorization_flow,
            signing_kp=CertificateKeyPair.objects.first(),
        )
        provider.property_mappings.set(SAMLPropertyMapping.objects.all())
        provider.save()
        Application.objects.create(
            name="SAML", slug="passbook-saml", provider=provider,
        )
        self.container = self.setup_client(provider)
        self.driver.get(
            self.url(
                "passbook_providers_saml:sso-init",
                application_slug=provider.application.slug,
            )
        )
        self.driver.find_element(By.ID, "id_uid_field").click()
        self.driver.find_element(By.ID, "id_uid_field").send_keys(USER().username)
        self.driver.find_element(By.ID, "id_uid_field").send_keys(Keys.ENTER)
        self.driver.find_element(By.ID, "id_password").send_keys(USER().username)
        self.driver.find_element(By.ID, "id_password").send_keys(Keys.ENTER)
        self.wait_for_url("http://localhost:9009/")
        self.assertEqual(
            self.driver.find_element(By.XPATH, "/html/body/pre").text,
            f"Hello, {USER().name}!",
        )

    def test_sp_initiated_denied(self):
        """test SAML Provider flow SP-initiated flow (Policy denies access)"""
        # Bootstrap all needed objects
        authorization_flow = Flow.objects.get(
            slug="default-provider-authorization-implicit-consent"
        )
        negative_policy = ExpressionPolicy.objects.create(
            name="negative-static", expression="return False"
        )
        provider: SAMLProvider = SAMLProvider.objects.create(
            name="saml-test",
            acs_url="http://localhost:9009/saml/acs",
            audience="passbook-e2e",
            issuer="passbook-e2e",
            sp_binding=SAMLBindings.POST,
            authorization_flow=authorization_flow,
            signing_kp=CertificateKeyPair.objects.first(),
        )
        provider.property_mappings.set(SAMLPropertyMapping.objects.all())
        provider.save()
        app = Application.objects.create(
            name="SAML", slug="passbook-saml", provider=provider,
        )
        PolicyBinding.objects.create(target=app, policy=negative_policy, order=0)
        self.container = self.setup_client(provider)
        self.driver.get("http://localhost:9009/")
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
