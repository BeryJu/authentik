"""passbook Policy administration"""
from typing import Any, Dict

from django.contrib import messages
from django.contrib.auth.mixins import LoginRequiredMixin
from django.contrib.auth.mixins import (
    PermissionRequiredMixin as DjangoPermissionRequiredMixin,
)
from django.contrib.messages.views import SuccessMessageMixin
from django.db.models import QuerySet
from django.forms import Form
from django.http import Http404, HttpRequest, HttpResponse
from django.urls import reverse_lazy
from django.utils.translation import ugettext as _
from django.views.generic import DeleteView, FormView, ListView, UpdateView
from django.views.generic.detail import DetailView
from guardian.mixins import PermissionListMixin, PermissionRequiredMixin

from passbook.admin.forms.policies import PolicyTestForm
from passbook.lib.utils.reflection import all_subclasses, path_to_class
from passbook.lib.views import CreateAssignPermView
from passbook.policies.models import Policy, PolicyBinding
from passbook.policies.process import PolicyProcess, PolicyRequest


class PolicyListView(LoginRequiredMixin, PermissionListMixin, ListView):
    """Show list of all policies"""

    model = Policy
    permission_required = "passbook_policies.view_policy"
    paginate_by = 10
    ordering = "name"
    template_name = "administration/policy/list.html"

    def get_context_data(self, **kwargs: Any) -> Dict[str, Any]:
        kwargs["types"] = {x.__name__: x for x in all_subclasses(Policy)}
        return super().get_context_data(**kwargs)

    def get_queryset(self) -> QuerySet:
        return super().get_queryset().select_subclasses()


class PolicyCreateView(
    SuccessMessageMixin,
    LoginRequiredMixin,
    DjangoPermissionRequiredMixin,
    CreateAssignPermView,
):
    """Create new Policy"""

    model = Policy
    permission_required = "passbook_policies.add_policy"

    template_name = "generic/create.html"
    success_url = reverse_lazy("passbook_admin:policies")
    success_message = _("Successfully created Policy")

    def get_context_data(self, **kwargs: Any) -> Dict[str, Any]:
        kwargs = super().get_context_data(**kwargs)
        form_cls = self.get_form_class()
        if hasattr(form_cls, "template_name"):
            kwargs["base_template"] = form_cls.template_name
        return kwargs

    def get_form_class(self) -> Form:
        policy_type = self.request.GET.get("type")
        try:
            model = next(x for x in all_subclasses(Policy) if x.__name__ == policy_type)
        except StopIteration as exc:
            raise Http404 from exc
        return path_to_class(model.form)


class PolicyUpdateView(
    SuccessMessageMixin, LoginRequiredMixin, PermissionRequiredMixin, UpdateView
):
    """Update policy"""

    model = Policy
    permission_required = "passbook_policies.change_policy"

    template_name = "generic/update.html"
    success_url = reverse_lazy("passbook_admin:policies")
    success_message = _("Successfully updated Policy")

    def get_context_data(self, **kwargs: Any) -> Dict[str, Any]:
        kwargs = super().get_context_data(**kwargs)
        form_cls = self.get_form_class()
        if hasattr(form_cls, "template_name"):
            kwargs["base_template"] = form_cls.template_name
        return kwargs

    def get_form_class(self) -> Form:
        form_class_path = self.get_object().form
        form_class = path_to_class(form_class_path)
        return form_class

    def get_object(self, queryset=None) -> Policy:
        return (
            Policy.objects.filter(pk=self.kwargs.get("pk")).select_subclasses().first()
        )


class PolicyDeleteView(
    SuccessMessageMixin, LoginRequiredMixin, PermissionRequiredMixin, DeleteView
):
    """Delete policy"""

    model = Policy
    permission_required = "passbook_policies.delete_policy"

    template_name = "generic/delete.html"
    success_url = reverse_lazy("passbook_admin:policies")
    success_message = _("Successfully deleted Policy")

    def get_object(self, queryset=None) -> Policy:
        return (
            Policy.objects.filter(pk=self.kwargs.get("pk")).select_subclasses().first()
        )

    def delete(self, request: HttpRequest, *args, **kwargs) -> HttpResponse:
        messages.success(self.request, self.success_message)
        return super().delete(request, *args, **kwargs)


class PolicyTestView(LoginRequiredMixin, DetailView, PermissionRequiredMixin, FormView):
    """View to test policy(s)"""

    model = Policy
    form_class = PolicyTestForm
    permission_required = "passbook_policies.view_policy"
    template_name = "administration/policy/test.html"
    object = None

    def get_object(self, queryset=None) -> QuerySet:
        return (
            Policy.objects.filter(pk=self.kwargs.get("pk")).select_subclasses().first()
        )

    def get_context_data(self, **kwargs: Any) -> Dict[str, Any]:
        kwargs["policy"] = self.get_object()
        return super().get_context_data(**kwargs)

    def post(self, *args, **kwargs) -> HttpResponse:
        self.object = self.get_object()
        return super().post(*args, **kwargs)

    def form_valid(self, form: PolicyTestForm) -> HttpResponse:
        policy = self.get_object()
        user = form.cleaned_data.get("user")

        p_request = PolicyRequest(user)
        p_request.http_request = self.request
        p_request.context = form.cleaned_data

        proc = PolicyProcess(PolicyBinding(policy=policy), p_request, None)
        result = proc.execute()
        if result:
            messages.success(self.request, _("User successfully passed policy."))
        else:
            messages.error(self.request, _("User didn't pass policy."))
        return self.render_to_response(self.get_context_data(form=form, result=result))
