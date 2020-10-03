"""passbook Group administration"""
from django.contrib.auth.mixins import LoginRequiredMixin
from django.contrib.auth.mixins import (
    PermissionRequiredMixin as DjangoPermissionRequiredMixin,
)
from django.contrib.messages.views import SuccessMessageMixin
from django.urls import reverse_lazy
from django.utils.translation import gettext as _
from django.views.generic import ListView, UpdateView
from guardian.mixins import PermissionListMixin, PermissionRequiredMixin

from passbook.admin.views.utils import (
    BackSuccessUrlMixin,
    DeleteMessageView,
    SearchListMixin,
    UserPaginateListMixin,
)
from passbook.core.forms.groups import GroupForm
from passbook.core.models import Group
from passbook.lib.views import CreateAssignPermView


class GroupListView(
    LoginRequiredMixin,
    PermissionListMixin,
    UserPaginateListMixin,
    SearchListMixin,
    ListView,
):
    """Show list of all groups"""

    model = Group
    permission_required = "passbook_core.view_group"
    ordering = "name"
    template_name = "administration/group/list.html"
    search_fields = ["name", "attributes"]


class GroupCreateView(
    SuccessMessageMixin,
    BackSuccessUrlMixin,
    LoginRequiredMixin,
    DjangoPermissionRequiredMixin,
    CreateAssignPermView,
):
    """Create new Group"""

    model = Group
    form_class = GroupForm
    permission_required = "passbook_core.add_group"

    template_name = "generic/create.html"
    success_url = reverse_lazy("passbook_admin:groups")
    success_message = _("Successfully created Group")


class GroupUpdateView(
    SuccessMessageMixin,
    BackSuccessUrlMixin,
    LoginRequiredMixin,
    PermissionRequiredMixin,
    UpdateView,
):
    """Update group"""

    model = Group
    form_class = GroupForm
    permission_required = "passbook_core.change_group"

    template_name = "generic/update.html"
    success_url = reverse_lazy("passbook_admin:groups")
    success_message = _("Successfully updated Group")


class GroupDeleteView(LoginRequiredMixin, PermissionRequiredMixin, DeleteMessageView):
    """Delete group"""

    model = Group
    permission_required = "passbook_flows.delete_group"

    template_name = "generic/delete.html"
    success_url = reverse_lazy("passbook_admin:groups")
    success_message = _("Successfully deleted Group")
