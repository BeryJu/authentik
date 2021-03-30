"""Discord OAuth Views"""
from typing import Any

from authentik.sources.oauth.models import OAuthSource, UserOAuthSourceConnection
from authentik.sources.oauth.types.manager import MANAGER, RequestKind
from authentik.sources.oauth.views.callback import OAuthCallback
from authentik.sources.oauth.views.redirect import OAuthRedirect


@MANAGER.source(kind=RequestKind.REDIRECT, name="Discord")
class DiscordOAuthRedirect(OAuthRedirect):
    """Discord OAuth2 Redirect"""

    def get_additional_parameters(self, source):  # pragma: no cover
        return {
            "scope": "email identify",
        }


@MANAGER.source(kind=RequestKind.CALLBACK, name="Discord")
class DiscordOAuth2Callback(OAuthCallback):
    """Discord OAuth2 Callback"""

    def get_user_enroll_context(
        self,
        source: OAuthSource,
        access: UserOAuthSourceConnection,
        info: dict[str, Any],
    ) -> dict[str, Any]:
        return {
            "username": info.get("username"),
            "email": info.get("email", None),
            "name": info.get("username"),
        }
