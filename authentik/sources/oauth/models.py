"""OAuth Client models"""
from typing import Optional, Type

from django.db import models
from django.forms import ModelForm
from django.urls import reverse, reverse_lazy
from django.utils.translation import gettext_lazy as _
from rest_framework.serializers import Serializer

from authentik.core.models import Source, UserSourceConnection
from authentik.core.types import UILoginButton


class OAuthSource(Source):
    """Login using a Generic OAuth provider."""

    provider_type = models.CharField(max_length=255)
    request_token_url = models.CharField(
        blank=True,
        max_length=255,
        verbose_name=_("Request Token URL"),
        help_text=_(
            "URL used to request the initial token. This URL is only required for OAuth 1."
        ),
    )
    authorization_url = models.CharField(
        max_length=255,
        verbose_name=_("Authorization URL"),
        help_text=_("URL the user is redirect to to conest the flow."),
    )
    access_token_url = models.CharField(
        max_length=255,
        verbose_name=_("Access Token URL"),
        help_text=_("URL used by authentik to retrive tokens."),
    )
    profile_url = models.CharField(
        max_length=255,
        verbose_name=_("Profile URL"),
        help_text=_("URL used by authentik to get user information."),
    )
    consumer_key = models.TextField()
    consumer_secret = models.TextField()

    @property
    def form(self) -> Type[ModelForm]:
        from authentik.sources.oauth.forms import OAuthSourceForm

        return OAuthSourceForm

    @property
    def serializer(self) -> Type[Serializer]:
        from authentik.sources.oauth.api import OAuthSourceSerializer

        return OAuthSourceSerializer

    @property
    def ui_login_button(self) -> UILoginButton:
        return UILoginButton(
            url=reverse_lazy(
                "authentik_sources_oauth:oauth-client-login",
                kwargs={"source_slug": self.slug},
            ),
            icon_path=f"authentik/sources/{self.provider_type}.svg",
            name=self.name,
        )

    @property
    def ui_additional_info(self) -> str:
        url = reverse_lazy(
            "authentik_sources_oauth:oauth-client-callback",
            kwargs={"source_slug": self.slug},
        )
        return f"Callback URL: <pre>{url}</pre>"

    @property
    def ui_user_settings(self) -> Optional[str]:
        view_name = "authentik_sources_oauth:oauth-client-user"
        return reverse(view_name, kwargs={"source_slug": self.slug})

    def __str__(self) -> str:
        return f"OAuth Source {self.name}"

    class Meta:

        verbose_name = _("Generic OAuth Source")
        verbose_name_plural = _("Generic OAuth Sources")


class GitHubOAuthSource(OAuthSource):
    """Social Login using GitHub.com or a GitHub-Enterprise Instance."""

    @property
    def form(self) -> Type[ModelForm]:
        from authentik.sources.oauth.forms import GitHubOAuthSourceForm

        return GitHubOAuthSourceForm

    class Meta:

        abstract = True
        verbose_name = _("GitHub OAuth Source")
        verbose_name_plural = _("GitHub OAuth Sources")


class TwitterOAuthSource(OAuthSource):
    """Social Login using Twitter.com"""

    @property
    def form(self) -> Type[ModelForm]:
        from authentik.sources.oauth.forms import TwitterOAuthSourceForm

        return TwitterOAuthSourceForm

    class Meta:

        abstract = True
        verbose_name = _("Twitter OAuth Source")
        verbose_name_plural = _("Twitter OAuth Sources")


class FacebookOAuthSource(OAuthSource):
    """Social Login using Facebook.com."""

    @property
    def form(self) -> Type[ModelForm]:
        from authentik.sources.oauth.forms import FacebookOAuthSourceForm

        return FacebookOAuthSourceForm

    class Meta:

        abstract = True
        verbose_name = _("Facebook OAuth Source")
        verbose_name_plural = _("Facebook OAuth Sources")


class DiscordOAuthSource(OAuthSource):
    """Social Login using Discord."""

    @property
    def form(self) -> Type[ModelForm]:
        from authentik.sources.oauth.forms import DiscordOAuthSourceForm

        return DiscordOAuthSourceForm

    class Meta:

        abstract = True
        verbose_name = _("Discord OAuth Source")
        verbose_name_plural = _("Discord OAuth Sources")


class GoogleOAuthSource(OAuthSource):
    """Social Login using Google or Gsuite."""

    @property
    def form(self) -> Type[ModelForm]:
        from authentik.sources.oauth.forms import GoogleOAuthSourceForm

        return GoogleOAuthSourceForm

    class Meta:

        abstract = True
        verbose_name = _("Google OAuth Source")
        verbose_name_plural = _("Google OAuth Sources")


class AzureADOAuthSource(OAuthSource):
    """Social Login using Azure AD."""

    @property
    def form(self) -> Type[ModelForm]:
        from authentik.sources.oauth.forms import AzureADOAuthSourceForm

        return AzureADOAuthSourceForm

    class Meta:

        abstract = True
        verbose_name = _("Azure AD OAuth Source")
        verbose_name_plural = _("Azure AD OAuth Sources")


class OpenIDOAuthSource(OAuthSource):
    """Login using a Generic OpenID-Connect compliant provider."""

    @property
    def form(self) -> Type[ModelForm]:
        from authentik.sources.oauth.forms import OAuthSourceForm

        return OAuthSourceForm

    class Meta:

        abstract = True
        verbose_name = _("OpenID OAuth Source")
        verbose_name_plural = _("OpenID OAuth Sources")


class UserOAuthSourceConnection(UserSourceConnection):
    """Authorized remote OAuth provider."""

    identifier = models.CharField(max_length=255)
    access_token = models.TextField(blank=True, null=True, default=None)

    def save(self, *args, **kwargs):
        self.access_token = self.access_token or None
        super().save(*args, **kwargs)

    class Meta:

        verbose_name = _("User OAuth Source Connection")
        verbose_name_plural = _("User OAuth Source Connections")
