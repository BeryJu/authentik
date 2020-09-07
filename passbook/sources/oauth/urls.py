"""passbook OAuth source urls"""

from django.urls import path

from passbook.sources.oauth.types.manager import RequestKind
from passbook.sources.oauth.views.dispatcher import DispatcherView
from passbook.sources.oauth.views.user import DisconnectView, UserSettingsView

urlpatterns = [
    path(
        "login/<slug:source_slug>/",
        DispatcherView.as_view(kind=RequestKind.redirect),
        name="oauth-client-login",
    ),
    path(
        "callback/<slug:source_slug>/",
        DispatcherView.as_view(kind=RequestKind.callback),
        name="oauth-client-callback",
    ),
    path(
        "user/<slug:source_slug>/",
        UserSettingsView.as_view(),
        name="oauth-client-user",
    ),
    path(
        "user/<slug:source_slug>/disconnect/",
        DisconnectView.as_view(),
        name="oauth-client-disconnect",
    ),
]
