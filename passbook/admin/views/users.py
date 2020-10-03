"""passbook User administration"""
from django.contrib import messages
from django.contrib.auth.mixins import LoginRequiredMixin
from django.contrib.auth.mixins import (
    PermissionRequiredMixin as DjangoPermissionRequiredMixin,
)
from django.contrib.messages.views import SuccessMessageMixin
from django.http import HttpRequest, HttpResponse
from django.shortcuts import redirect
from django.urls import reverse, reverse_lazy
from django.utils.http import urlencode
from django.utils.translation import gettext as _
from django.views.generic import DetailView, ListView, UpdateView
from guardian.mixins import (
    PermissionListMixin,
    PermissionRequiredMixin,
    get_anonymous_user,
)

from passbook.admin.forms.users import UserForm
from passbook.admin.views.utils import (
    BackSuccessUrlMixin,
    DeleteMessageView,
    SearchListMixin,
    UserPaginateListMixin,
)
from passbook.core.models import Token, User
from passbook.lib.views import CreateAssignPermView


class UserListView(
    LoginRequiredMixin,
    PermissionListMixin,
    UserPaginateListMixin,
    SearchListMixin,
    ListView,
):
    """Show list of all users"""

    model = User
    permission_required = "passbook_core.view_user"
    ordering = "username"
    template_name = "administration/user/list.html"
    search_fields = ["username", "name", "attributes"]

    def get_queryset(self):
        return super().get_queryset().exclude(pk=get_anonymous_user().pk)


class UserCreateView(
    SuccessMessageMixin,
    BackSuccessUrlMixin,
    LoginRequiredMixin,
    DjangoPermissionRequiredMixin,
    CreateAssignPermView,
):
    """Create user"""

    model = User
    form_class = UserForm
    permission_required = "passbook_core.add_user"

    template_name = "generic/create.html"
    success_url = reverse_lazy("passbook_admin:users")
    success_message = _("Successfully created User")


class UserUpdateView(
    SuccessMessageMixin,
    BackSuccessUrlMixin,
    LoginRequiredMixin,
    PermissionRequiredMixin,
    UpdateView,
):
    """Update user"""

    model = User
    form_class = UserForm
    permission_required = "passbook_core.change_user"

    # By default the object's name is user which is used by other checks
    context_object_name = "object"
    template_name = "generic/update.html"
    success_url = reverse_lazy("passbook_admin:users")
    success_message = _("Successfully updated User")


class UserDeleteView(LoginRequiredMixin, PermissionRequiredMixin, DeleteMessageView):
    """Delete user"""

    model = User
    permission_required = "passbook_core.delete_user"

    # By default the object's name is user which is used by other checks
    context_object_name = "object"
    template_name = "generic/delete.html"
    success_url = reverse_lazy("passbook_admin:users")
    success_message = _("Successfully deleted User")


class UserPasswordResetView(LoginRequiredMixin, PermissionRequiredMixin, DetailView):
    """Get Password reset link for user"""

    model = User
    permission_required = "passbook_core.reset_user_password"

    def get(self, request: HttpRequest, *args, **kwargs) -> HttpResponse:
        """Create token for user and return link"""
        super().get(request, *args, **kwargs)
        token = Token.objects.create(user=self.object)
        querystring = urlencode({"token": token.token_uuid})
        link = request.build_absolute_uri(
            reverse("passbook_flows:default-recovery") + f"?{querystring}"
        )
        messages.success(
            request, _("Password reset link: <pre>%(link)s</pre>" % {"link": link})
        )
        return redirect("passbook_admin:users")
