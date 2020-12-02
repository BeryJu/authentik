"""authentik Provider administration"""
from django.contrib.auth.mixins import LoginRequiredMixin
from django.contrib.auth.mixins import (
    PermissionRequiredMixin as DjangoPermissionRequiredMixin,
)
from django.contrib.messages.views import SuccessMessageMixin
from django.urls import reverse_lazy
from django.utils.translation import gettext as _
from guardian.mixins import PermissionListMixin, PermissionRequiredMixin

from authentik.admin.views.utils import (
    BackSuccessUrlMixin,
    DeleteMessageView,
    InheritanceCreateView,
    InheritanceListView,
    InheritanceUpdateView,
    SearchListMixin,
    UserPaginateListMixin,
)
from authentik.core.models import Provider


class ProviderListView(
    LoginRequiredMixin,
    PermissionListMixin,
    UserPaginateListMixin,
    SearchListMixin,
    InheritanceListView,
):
    """Show list of all providers"""

    model = Provider
    permission_required = "authentik_core.add_provider"
    template_name = "administration/provider/list.html"
    ordering = "pk"
    search_fields = ["pk", "name"]


class ProviderCreateView(
    SuccessMessageMixin,
    BackSuccessUrlMixin,
    LoginRequiredMixin,
    DjangoPermissionRequiredMixin,
    InheritanceCreateView,
):
    """Create new Provider"""

    model = Provider
    permission_required = "authentik_core.add_provider"

    template_name = "generic/create.html"
    success_url = reverse_lazy("authentik_admin:providers")
    success_message = _("Successfully created Provider")


class ProviderUpdateView(
    SuccessMessageMixin,
    BackSuccessUrlMixin,
    LoginRequiredMixin,
    PermissionRequiredMixin,
    InheritanceUpdateView,
):
    """Update provider"""

    model = Provider
    permission_required = "authentik_core.change_provider"

    template_name = "generic/update.html"
    success_url = reverse_lazy("authentik_admin:providers")
    success_message = _("Successfully updated Provider")


class ProviderDeleteView(
    LoginRequiredMixin, PermissionRequiredMixin, DeleteMessageView
):
    """Delete provider"""

    model = Provider
    permission_required = "authentik_core.delete_provider"

    template_name = "generic/delete.html"
    success_url = reverse_lazy("authentik_admin:providers")
    success_message = _("Successfully deleted Provider")
