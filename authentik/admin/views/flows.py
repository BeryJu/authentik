"""authentik Flow administration"""
from django.contrib import messages
from django.contrib.auth.mixins import LoginRequiredMixin
from django.contrib.auth.mixins import (
    PermissionRequiredMixin as DjangoPermissionRequiredMixin,
)
from django.contrib.messages.views import SuccessMessageMixin
from django.http import HttpRequest, HttpResponse
from django.urls import reverse_lazy
from django.utils.translation import gettext as _
from django.views.generic import DetailView, FormView, UpdateView
from guardian.mixins import PermissionRequiredMixin

from authentik.flows.exceptions import FlowNonApplicableException
from authentik.flows.forms import FlowForm, FlowImportForm
from authentik.flows.models import Flow
from authentik.flows.planner import PLAN_CONTEXT_PENDING_USER
from authentik.flows.transfer.importer import FlowImporter
from authentik.flows.views import SESSION_KEY_PLAN, FlowPlanner
from authentik.lib.utils.urls import redirect_with_qs
from authentik.lib.views import CreateAssignPermView, bad_request_message


class FlowCreateView(
    SuccessMessageMixin,
    LoginRequiredMixin,
    DjangoPermissionRequiredMixin,
    CreateAssignPermView,
):
    """Create new Flow"""

    model = Flow
    form_class = FlowForm
    permission_required = "authentik_flows.add_flow"

    template_name = "generic/create.html"
    success_url = reverse_lazy("authentik_core:if-admin")
    success_message = _("Successfully created Flow")


class FlowUpdateView(
    SuccessMessageMixin,
    LoginRequiredMixin,
    PermissionRequiredMixin,
    UpdateView,
):
    """Update flow"""

    model = Flow
    form_class = FlowForm
    permission_required = "authentik_flows.change_flow"

    template_name = "generic/update.html"
    success_url = reverse_lazy("authentik_core:if-admin")
    success_message = _("Successfully updated Flow")


class FlowDebugExecuteView(LoginRequiredMixin, PermissionRequiredMixin, DetailView):
    """Debug exectue flow, setting the current user as pending user"""

    model = Flow
    permission_required = "authentik_flows.view_flow"

    # pylint: disable=unused-argument
    def get(self, request: HttpRequest, pk: str) -> HttpResponse:
        """Debug exectue flow, setting the current user as pending user"""
        flow: Flow = self.get_object()
        planner = FlowPlanner(flow)
        planner.use_cache = False
        try:
            plan = planner.plan(self.request, {PLAN_CONTEXT_PENDING_USER: request.user})
            self.request.session[SESSION_KEY_PLAN] = plan
        except FlowNonApplicableException as exc:
            return bad_request_message(
                request,
                _(
                    "Flow not applicable to current user/request: %(messages)s"
                    % {"messages": str(exc)}
                ),
            )
        return redirect_with_qs(
            "authentik_core:if-flow",
            self.request.GET,
            flow_slug=flow.slug,
        )


class FlowImportView(LoginRequiredMixin, FormView):
    """Import flow from JSON Export; only allowed for superusers
    as these flows can contain python code"""

    form_class = FlowImportForm
    template_name = "administration/flow/import.html"
    success_url = reverse_lazy("authentik_core:if-admin")

    def dispatch(self, request, *args, **kwargs):
        if not request.user.is_superuser:
            return self.handle_no_permission()
        return super().dispatch(request, *args, **kwargs)

    def form_valid(self, form: FlowImportForm) -> HttpResponse:
        importer = FlowImporter(form.cleaned_data["flow"].read().decode())
        successful = importer.apply()
        if not successful:
            messages.error(self.request, _("Failed to import flow."))
        else:
            messages.success(self.request, _("Successfully imported flow."))
        return super().form_valid(form)
