"""Google OAuth Views"""
from typing import Any

from authentik.sources.oauth.models import OAuthSource, UserOAuthSourceConnection
from authentik.sources.oauth.types.manager import MANAGER, RequestKind
from authentik.sources.oauth.views.callback import OAuthCallback
from authentik.sources.oauth.views.redirect import OAuthRedirect


@MANAGER.source(kind=RequestKind.REDIRECT, name="Google")
class GoogleOAuthRedirect(OAuthRedirect):
    """Google OAuth2 Redirect"""

    def get_additional_parameters(self, source):  # pragma: no cover
        return {
            "scope": "email profile",
        }


@MANAGER.source(kind=RequestKind.CALLBACK, name="Google")
class GoogleOAuth2Callback(OAuthCallback):
    """Google OAuth2 Callback"""

    def get_user_enroll_context(
        self,
        source: OAuthSource,
        access: UserOAuthSourceConnection,
        info: dict[str, Any],
    ) -> dict[str, Any]:
        return {
            "username": info.get("email"),
            "email": info.get("email"),
            "name": info.get("name"),
        }
