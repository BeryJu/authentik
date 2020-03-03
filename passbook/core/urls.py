"""passbook URL Configuration"""
from django.urls import path
from structlog import get_logger

from passbook.core.views import authentication, overview, user
from passbook.flows.executor.http import FactorPermissionDeniedView, HttpExecutorView

LOGGER = get_logger()

urlpatterns = [
    # Authentication views
    path("auth/login/", authentication.LoginView.as_view(), name="auth-login"),
    path("auth/logout/", authentication.LogoutView.as_view(), name="auth-logout"),
    path("auth/sign_up/", authentication.SignUpView.as_view(), name="auth-sign-up"),
    path(
        "auth/sign_up/<uuid:nonce>/confirm/",
        authentication.SignUpConfirmView.as_view(),
        name="auth-sign-up-confirm",
    ),
    path(
        "auth/process/denied/",
        FactorPermissionDeniedView.as_view(),
        name="auth-denied",
    ),
    path(
        "auth/password/reset/<uuid:nonce>/",
        authentication.PasswordResetView.as_view(),
        name="auth-password-reset",
    ),
    path("flows/execute/", HttpExecutorView.as_view(), name="flows-execute"),
    # User views
    path("_/user/", user.UserSettingsView.as_view(), name="user-settings"),
    path("_/user/delete/", user.UserDeleteView.as_view(), name="user-delete"),
    path(
        "_/user/change_password/",
        user.UserChangePasswordView.as_view(),
        name="user-change-password",
    ),
    # Overview
    path("", overview.OverviewView.as_view(), name="overview"),
]
