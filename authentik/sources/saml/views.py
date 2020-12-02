"""saml sp views"""
from django.contrib.auth import logout
from django.contrib.auth.mixins import LoginRequiredMixin
from django.http import Http404, HttpRequest, HttpResponse
from django.shortcuts import get_object_or_404, redirect, render
from django.utils.decorators import method_decorator
from django.utils.http import urlencode
from django.utils.translation import gettext_lazy as _
from django.views import View
from django.views.decorators.csrf import csrf_exempt
from xmlsec import VerificationError

from authentik.lib.views import bad_request_message
from authentik.providers.saml.utils.encoding import nice64
from authentik.sources.saml.exceptions import (
    MissingSAMLResponse,
    UnsupportedNameIDFormat,
)
from authentik.sources.saml.models import SAMLBindingTypes, SAMLSource
from authentik.sources.saml.processors.metadata import MetadataProcessor
from authentik.sources.saml.processors.request import RequestProcessor
from authentik.sources.saml.processors.response import ResponseProcessor


class InitiateView(View):
    """Get the Form with SAML Request, which sends us to the IDP"""

    def get(self, request: HttpRequest, source_slug: str) -> HttpResponse:
        """Replies with an XHTML SSO Request."""
        source: SAMLSource = get_object_or_404(SAMLSource, slug=source_slug)
        if not source.enabled:
            raise Http404
        relay_state = request.GET.get("next", "")
        auth_n_req = RequestProcessor(source, request, relay_state)
        # If the source is configured for Redirect bindings, we can just redirect there
        if source.binding_type == SAMLBindingTypes.Redirect:
            url_args = urlencode(auth_n_req.build_auth_n_detached())
            return redirect(f"{source.sso_url}?{url_args}")
        # As POST Binding we show a form
        saml_request = nice64(auth_n_req.build_auth_n())
        if source.binding_type == SAMLBindingTypes.POST:
            return render(
                request,
                "saml/sp/login.html",
                {
                    "request_url": source.sso_url,
                    "request": saml_request,
                    "relay_state": relay_state,
                    "source": source,
                },
            )
        # Or an auto-submit form
        if source.binding_type == SAMLBindingTypes.POST_AUTO:
            return render(
                request,
                "generic/autosubmit_form_full.html",
                {
                    "title": _("Redirecting to %(app)s..." % {"app": source.name}),
                    "attrs": {"SAMLRequest": saml_request, "RelayState": relay_state},
                    "url": source.sso_url,
                },
            )
        raise Http404


@method_decorator(csrf_exempt, name="dispatch")
class ACSView(View):
    """AssertionConsumerService, consume assertion and log user in"""

    def post(self, request: HttpRequest, source_slug: str) -> HttpResponse:
        """Handles a POSTed SSO Assertion and logs the user in."""
        source: SAMLSource = get_object_or_404(SAMLSource, slug=source_slug)
        if not source.enabled:
            raise Http404
        processor = ResponseProcessor(source)
        try:
            processor.parse(request)
        except MissingSAMLResponse as exc:
            return bad_request_message(request, str(exc))
        except VerificationError as exc:
            return bad_request_message(request, str(exc))

        try:
            return processor.prepare_flow(request)
        except UnsupportedNameIDFormat as exc:
            return bad_request_message(request, str(exc))


class SLOView(LoginRequiredMixin, View):
    """Single-Logout-View"""

    def dispatch(self, request: HttpRequest, source_slug: str) -> HttpResponse:
        """Log user out and redirect them to the IdP's SLO URL."""
        source: SAMLSource = get_object_or_404(SAMLSource, slug=source_slug)
        if not source.enabled:
            raise Http404
        logout(request)
        return redirect(source.slo_url)


class MetadataView(View):
    """Return XML Metadata for IDP"""

    def dispatch(self, request: HttpRequest, source_slug: str) -> HttpResponse:
        """Replies with the XML Metadata SPSSODescriptor."""
        source: SAMLSource = get_object_or_404(SAMLSource, slug=source_slug)
        metadata = MetadataProcessor(source, request).build_entity_descriptor()
        return HttpResponse(metadata, content_type="text/xml")
