"""authentik SAML IDP Views"""
from typing import Optional

from django.http import HttpRequest, HttpResponse
from django.shortcuts import get_object_or_404
from django.utils.decorators import method_decorator
from django.views.decorators.csrf import csrf_exempt
from structlog.stdlib import get_logger

from authentik.core.models import Application
from authentik.events.models import Event, EventAction
from authentik.flows.models import in_memory_stage
from authentik.flows.planner import (
    PLAN_CONTEXT_APPLICATION,
    PLAN_CONTEXT_SSO,
    FlowPlanner,
)
from authentik.flows.views import SESSION_KEY_PLAN
from authentik.lib.utils.urls import redirect_with_qs
from authentik.lib.views import bad_request_message
from authentik.policies.views import PolicyAccessView
from authentik.providers.saml.exceptions import CannotHandleAssertion
from authentik.providers.saml.models import SAMLProvider
from authentik.providers.saml.processors.request_parser import AuthNRequestParser
from authentik.providers.saml.views.flows import (
    REQUEST_KEY_RELAY_STATE,
    REQUEST_KEY_SAML_REQUEST,
    REQUEST_KEY_SAML_SIG_ALG,
    REQUEST_KEY_SAML_SIGNATURE,
    SESSION_KEY_AUTH_N_REQUEST,
    SAMLFlowFinalView,
)
from authentik.stages.consent.stage import PLAN_CONTEXT_CONSENT_TEMPLATE

LOGGER = get_logger()


class SAMLSSOView(PolicyAccessView):
    """ "SAML SSO Base View, which plans a flow and injects our final stage.
    Calls get/post handler."""

    def resolve_provider_application(self):
        self.application = get_object_or_404(
            Application, slug=self.kwargs["application_slug"]
        )
        self.provider: SAMLProvider = get_object_or_404(
            SAMLProvider, pk=self.application.provider_id
        )

    def check_saml_request(self) -> Optional[HttpRequest]:
        """Handler to verify the SAML Request. Must be implemented by a subclass"""
        raise NotImplementedError

    # pylint: disable=unused-argument
    def get(self, request: HttpRequest, application_slug: str) -> HttpResponse:
        """Verify the SAML Request, and if valid initiate the FlowPlanner for the application"""
        # Call the method handler, which checks the SAML
        # Request and returns a HTTP Response on error
        method_response = self.check_saml_request()
        if method_response:
            return method_response
        # Regardless, we start the planner and return to it
        planner = FlowPlanner(self.provider.authorization_flow)
        planner.allow_empty_flows = True
        plan = planner.plan(
            request,
            {
                PLAN_CONTEXT_SSO: True,
                PLAN_CONTEXT_APPLICATION: self.application,
                PLAN_CONTEXT_CONSENT_TEMPLATE: "providers/saml/consent.html",
            },
        )
        plan.append(in_memory_stage(SAMLFlowFinalView))
        request.session[SESSION_KEY_PLAN] = plan
        return redirect_with_qs(
            "authentik_flows:flow-executor-shell",
            request.GET,
            flow_slug=self.provider.authorization_flow.slug,
        )

    def post(self, request: HttpRequest, application_slug: str) -> HttpResponse:
        """GET and POST use the same handler, but we can't
        override .dispatch easily because PolicyAccessView's dispatch"""
        return self.get(request, application_slug)


class SAMLSSOBindingRedirectView(SAMLSSOView):
    """SAML Handler for SSO/Redirect bindings, which are sent via GET"""

    def check_saml_request(self) -> Optional[HttpRequest]:
        """Handle REDIRECT bindings"""
        if REQUEST_KEY_SAML_REQUEST not in self.request.GET:
            LOGGER.info("handle_saml_request: SAML payload missing")
            return bad_request_message(
                self.request, "The SAML request payload is missing."
            )

        try:
            auth_n_request = AuthNRequestParser(self.provider).parse_detached(
                self.request.GET[REQUEST_KEY_SAML_REQUEST],
                self.request.GET.get(REQUEST_KEY_RELAY_STATE),
                self.request.GET.get(REQUEST_KEY_SAML_SIGNATURE),
                self.request.GET.get(REQUEST_KEY_SAML_SIG_ALG),
            )
            self.request.session[SESSION_KEY_AUTH_N_REQUEST] = auth_n_request
        except CannotHandleAssertion as exc:
            Event.new(
                EventAction.CONFIGURATION_ERROR,
                provider=self.provider,
                message=str(exc),
            ).save()
            LOGGER.info(str(exc))
            return bad_request_message(self.request, str(exc))
        return None


@method_decorator(csrf_exempt, name="dispatch")
class SAMLSSOBindingPOSTView(SAMLSSOView):
    """SAML Handler for SSO/POST bindings"""

    def check_saml_request(self) -> Optional[HttpRequest]:
        """Handle POST bindings"""
        if REQUEST_KEY_SAML_REQUEST not in self.request.POST:
            LOGGER.info("check_saml_request: SAML payload missing")
            return bad_request_message(
                self.request, "The SAML request payload is missing."
            )

        try:
            auth_n_request = AuthNRequestParser(self.provider).parse(
                self.request.POST[REQUEST_KEY_SAML_REQUEST],
                self.request.POST.get(REQUEST_KEY_RELAY_STATE),
            )
            self.request.session[SESSION_KEY_AUTH_N_REQUEST] = auth_n_request
        except CannotHandleAssertion as exc:
            LOGGER.info(str(exc))
            return bad_request_message(self.request, str(exc))
        return None


class SAMLSSOBindingInitView(SAMLSSOView):
    """SAML Handler for for IdP Initiated login flows"""

    def check_saml_request(self) -> Optional[HttpRequest]:
        """Create SAML Response from scratch"""
        LOGGER.debug(
            "handle_saml_no_request: No SAML Request, using IdP-initiated flow."
        )
        auth_n_request = AuthNRequestParser(self.provider).idp_initiated()
        self.request.session[SESSION_KEY_AUTH_N_REQUEST] = auth_n_request