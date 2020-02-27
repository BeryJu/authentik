"""GitHub OAuth Views"""
from passbook.sources.oauth.types.manager import MANAGER, RequestKind
from passbook.sources.oauth.utils import user_get_or_create
from passbook.sources.oauth.views.core import OAuthCallback


@MANAGER.source(kind=RequestKind.callback, name="GitHub")
class GitHubOAuth2Callback(OAuthCallback):
    """GitHub OAuth2 Callback"""

    def get_or_create_user(self, source, access, info):
        user_data = {
            "username": info.get("login"),
            "email": info.get("email", ""),
            "name": info.get("name"),
            "password": None,
        }
        gh_user = user_get_or_create(**user_data)
        return gh_user
