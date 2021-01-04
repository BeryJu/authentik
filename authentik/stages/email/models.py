"""email stage models"""
from pathlib import Path
from typing import Type

from django.conf import settings
from django.core.mail import get_connection
from django.core.mail.backends.base import BaseEmailBackend
from django.db import models
from django.forms import ModelForm
from django.utils.translation import gettext as _
from django.views import View
from rest_framework.serializers import BaseSerializer

from authentik.flows.models import Stage


class EmailTemplates(models.TextChoices):
    """Templates used for rendering the Email"""

    PASSWORD_RESET = (
        "email/password_reset.html",
        _("Password Reset"),
    )  # nosec
    ACCOUNT_CONFIRM = (
        "email/account_confirmation.html",
        _("Account Confirmation"),
    )


def get_template_choices():
    """Get all available Email templates, including dynamically mounted ones.
    Directories are taken from TEMPLATES.DIR setting"""
    static_choices = EmailTemplates.choices

    dirs = [Path(x) for x in settings.TEMPLATES[0]["DIRS"]]
    for template_dir in dirs:
        if not template_dir.exists():
            continue
        for template in template_dir.glob("**/*.html"):
            path = str(template)
            static_choices.append((path, f"Custom Template: {path}"))
    return static_choices


class EmailStage(Stage):
    """Sends an Email to the user with a token to confirm their Email address."""

    use_global_settings = models.BooleanField(
        default=False,
        help_text=_(
            (
                "When enabled, global Email connection settings will be used and "
                "connection settings below will be ignored."
            )
        ),
    )

    host = models.TextField(default="localhost")
    port = models.IntegerField(default=25)
    username = models.TextField(default="", blank=True)
    password = models.TextField(default="", blank=True)
    use_tls = models.BooleanField(default=False)
    use_ssl = models.BooleanField(default=False)
    timeout = models.IntegerField(default=10)
    from_address = models.EmailField(default="system@authentik.local")

    token_expiry = models.IntegerField(
        default=30, help_text=_("Time in minutes the token sent is valid.")
    )
    subject = models.TextField(default="authentik")
    template = models.TextField(
        choices=get_template_choices(), default=EmailTemplates.PASSWORD_RESET
    )

    @property
    def serializer(self) -> BaseSerializer:
        from authentik.stages.email.api import EmailStageSerializer

        return EmailStageSerializer

    @property
    def type(self) -> Type[View]:
        from authentik.stages.email.stage import EmailStageView

        return EmailStageView

    @property
    def form(self) -> Type[ModelForm]:
        from authentik.stages.email.forms import EmailStageForm

        return EmailStageForm

    @property
    def backend(self) -> BaseEmailBackend:
        """Get fully configured Email Backend instance"""
        if self.use_global_settings:
            return get_connection()
        return get_connection(
            host=self.host,
            port=self.port,
            username=self.username,
            password=self.password,
            use_tls=self.use_tls,
            use_ssl=self.use_ssl,
            timeout=self.timeout,
        )

    def __str__(self):
        return f"Email Stage {self.name}"

    class Meta:

        verbose_name = _("Email Stage")
        verbose_name_plural = _("Email Stages")
