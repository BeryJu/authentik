"""email tests"""
from unittest.mock import MagicMock, patch

from django.core import mail
from django.shortcuts import reverse
from django.test import Client, TestCase
from django.utils.encoding import force_text

from passbook.core.models import Token, User
from passbook.flows.markers import StageMarker
from passbook.flows.models import Flow, FlowDesignation, FlowStageBinding
from passbook.flows.planner import PLAN_CONTEXT_PENDING_USER, FlowPlan
from passbook.flows.views import SESSION_KEY_PLAN
from passbook.stages.email.models import EmailStage
from passbook.stages.email.stage import QS_KEY_TOKEN


class TestEmailStage(TestCase):
    """Email tests"""

    def setUp(self):
        super().setUp()
        self.user = User.objects.create_user(
            username="unittest", email="test@beryju.org"
        )
        self.client = Client()

        self.flow = Flow.objects.create(
            name="test-email",
            slug="test-email",
            designation=FlowDesignation.AUTHENTICATION,
        )
        self.stage = EmailStage.objects.create(name="email",)
        FlowStageBinding.objects.create(flow=self.flow, stage=self.stage, order=2)

    def test_rendering(self):
        """Test with pending user"""
        plan = FlowPlan(
            flow_pk=self.flow.pk.hex, stages=[self.stage], markers=[StageMarker()]
        )
        plan.context[PLAN_CONTEXT_PENDING_USER] = self.user
        session = self.client.session
        session[SESSION_KEY_PLAN] = plan
        session.save()

        url = reverse(
            "passbook_flows:flow-executor", kwargs={"flow_slug": self.flow.slug}
        )
        response = self.client.get(url)
        self.assertEqual(response.status_code, 200)

    def test_without_user(self):
        """Test without pending user"""
        plan = FlowPlan(
            flow_pk=self.flow.pk.hex, stages=[self.stage], markers=[StageMarker()]
        )
        session = self.client.session
        session[SESSION_KEY_PLAN] = plan
        session.save()

        url = reverse(
            "passbook_flows:flow-executor", kwargs={"flow_slug": self.flow.slug}
        )
        response = self.client.get(url)
        self.assertEqual(response.status_code, 200)

    def test_pending_user(self):
        """Test with pending user"""
        plan = FlowPlan(
            flow_pk=self.flow.pk.hex, stages=[self.stage], markers=[StageMarker()]
        )
        plan.context[PLAN_CONTEXT_PENDING_USER] = self.user
        session = self.client.session
        session[SESSION_KEY_PLAN] = plan
        session.save()

        url = reverse(
            "passbook_flows:flow-executor", kwargs={"flow_slug": self.flow.slug}
        )
        with self.settings(
            EMAIL_BACKEND="django.core.mail.backends.locmem.EmailBackend"
        ):
            response = self.client.post(url)
            self.assertEqual(response.status_code, 200)
            self.assertEqual(len(mail.outbox), 1)
            self.assertEqual(mail.outbox[0].subject, "passbook - Password Recovery")

    def test_token(self):
        """Test with token"""
        # Make sure token exists
        self.test_pending_user()
        plan = FlowPlan(
            flow_pk=self.flow.pk.hex, stages=[self.stage], markers=[StageMarker()]
        )
        session = self.client.session
        session[SESSION_KEY_PLAN] = plan
        session.save()

        with patch("passbook.flows.views.FlowExecutorView.cancel", MagicMock()):
            url = reverse(
                "passbook_flows:flow-executor", kwargs={"flow_slug": self.flow.slug}
            )
            token = Token.objects.get(user=self.user)
            url += f"?{QS_KEY_TOKEN}={token.pk.hex}"
            response = self.client.get(url)

            self.assertEqual(response.status_code, 200)
            self.assertJSONEqual(
                force_text(response.content),
                {"type": "redirect", "to": reverse("passbook_core:overview")},
            )

            session = self.client.session
            plan: FlowPlan = session[SESSION_KEY_PLAN]
            self.assertEqual(plan.context[PLAN_CONTEXT_PENDING_USER], self.user)
