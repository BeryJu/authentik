"""OTP Time-based models"""
from typing import Optional, Type

from django.db import models
from django.forms import ModelForm
from django.shortcuts import reverse
from django.utils.translation import gettext_lazy as _
from django.views import View
from rest_framework.serializers import BaseSerializer

from authentik.flows.models import ConfigurableStage, Stage


class TOTPDigits(models.IntegerChoices):
    """OTP Time Digits"""

    SIX = 6, _("6 digits, widely compatible")
    EIGHT = 8, _("8 digits, not compatible with apps like Google Authenticator")


class OTPTimeStage(ConfigurableStage, Stage):
    """Enroll a user's device into Time-based OTP."""

    digits = models.IntegerField(choices=TOTPDigits.choices)

    @property
    def serializer(self) -> BaseSerializer:
        from authentik.stages.otp_time.api import OTPTimeStageSerializer

        return OTPTimeStageSerializer

    @property
    def type(self) -> Type[View]:
        from authentik.stages.otp_time.stage import OTPTimeStageView

        return OTPTimeStageView

    @property
    def form(self) -> Type[ModelForm]:
        from authentik.stages.otp_time.forms import OTPTimeStageForm

        return OTPTimeStageForm

    @property
    def ui_user_settings(self) -> Optional[str]:
        return reverse(
            "authentik_stages_otp_time:user-settings",
            kwargs={"stage_uuid": self.stage_uuid},
        )

    def __str__(self) -> str:
        return f"OTP Time (TOTP) Stage {self.name}"

    class Meta:

        verbose_name = _("OTP Time (TOTP) Setup Stage")
        verbose_name_plural = _("OTP Time (TOTP) Setup Stages")
