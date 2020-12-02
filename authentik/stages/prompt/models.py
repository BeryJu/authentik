"""prompt models"""
from typing import Type
from uuid import uuid4

from django import forms
from django.db import models
from django.forms import ModelForm
from django.utils.translation import gettext_lazy as _
from django.views import View
from rest_framework.serializers import BaseSerializer

from authentik.flows.models import Stage
from authentik.lib.models import SerializerModel
from authentik.policies.models import Policy
from authentik.stages.prompt.widgets import HorizontalRuleWidget, StaticTextWidget


class FieldTypes(models.TextChoices):
    """Field types an Prompt can be"""

    # Simple text field
    TEXT = "text", _("Text: Simple Text input")
    # Same as text, but has autocomplete for password managers
    USERNAME = (
        "username",
        _(
            (
                "Username: Same as Text input, but checks for "
                "and prevents duplicate usernames."
            )
        ),
    )
    EMAIL = "email", _("Email: Text field with Email type.")
    PASSWORD = (
        "password",  # noqa # nosec
        _(
            (
                "Password: Masked input, password is validated against sources. Policies still "
                "have to be applied to this Stage. If two of these are used in the same stage, "
                "they are ensured to be identical."
            )
        ),
    )
    NUMBER = "number"
    CHECKBOX = "checkbox"
    DATE = "data"
    DATE_TIME = "data-time"

    SEPARATOR = "separator", _("Separator: Static Separator Line")
    HIDDEN = "hidden", _("Hidden: Hidden field, can be used to insert data into form.")
    STATIC = "static", _("Static: Static value, displayed as-is.")


class Prompt(SerializerModel):
    """Single Prompt, part of a prompt stage."""

    prompt_uuid = models.UUIDField(primary_key=True, editable=False, default=uuid4)

    field_key = models.SlugField(
        help_text=_("Name of the form field, also used to store the value")
    )
    label = models.TextField()
    type = models.CharField(max_length=100, choices=FieldTypes.choices)
    required = models.BooleanField(default=True)
    placeholder = models.TextField(blank=True)

    order = models.IntegerField(default=0)

    @property
    def serializer(self) -> BaseSerializer:
        from authentik.stages.prompt.api import PromptSerializer

        return PromptSerializer

    @property
    def field(self):
        """Return instantiated form input field"""
        attrs = {"placeholder": _(self.placeholder)}
        field_class = forms.CharField
        widget = forms.TextInput(attrs=attrs)
        kwargs = {
            "label": _(self.label),
            "required": self.required,
        }
        if self.type == FieldTypes.EMAIL:
            field_class = forms.EmailField
        if self.type == FieldTypes.USERNAME:
            attrs["autocomplete"] = "username"
        if self.type == FieldTypes.PASSWORD:
            widget = forms.PasswordInput(attrs=attrs)
            attrs["autocomplete"] = "new-password"
        if self.type == FieldTypes.NUMBER:
            field_class = forms.IntegerField
            widget = forms.NumberInput(attrs=attrs)
        if self.type == FieldTypes.HIDDEN:
            widget = forms.HiddenInput(attrs=attrs)
            kwargs["required"] = False
            kwargs["initial"] = self.placeholder
        if self.type == FieldTypes.CHECKBOX:
            field_class = forms.BooleanField
            kwargs["required"] = False
        if self.type == FieldTypes.DATE:
            attrs["type"] = "date"
            widget = forms.DateInput(attrs=attrs)
        if self.type == FieldTypes.DATE_TIME:
            attrs["type"] = "datetime-local"
            widget = forms.DateTimeInput(attrs=attrs)
        if self.type == FieldTypes.STATIC:
            widget = StaticTextWidget(attrs=attrs)
            kwargs["initial"] = self.placeholder
            kwargs["required"] = False
            kwargs["label"] = ""
        if self.type == FieldTypes.SEPARATOR:
            widget = HorizontalRuleWidget(attrs=attrs)
            kwargs["required"] = False
            kwargs["label"] = ""

        kwargs["widget"] = widget
        return field_class(**kwargs)

    def save(self, *args, **kwargs):
        if self.type not in FieldTypes:
            raise ValueError
        return super().save(*args, **kwargs)

    def __str__(self):
        return f"Prompt '{self.field_key}' type={self.type}"

    class Meta:

        verbose_name = _("Prompt")
        verbose_name_plural = _("Prompts")


class PromptStage(Stage):
    """Define arbitrary prompts for the user."""

    fields = models.ManyToManyField(Prompt)

    validation_policies = models.ManyToManyField(Policy, blank=True)

    @property
    def serializer(self) -> BaseSerializer:
        from authentik.stages.prompt.api import PromptStageSerializer

        return PromptStageSerializer

    @property
    def type(self) -> Type[View]:
        from authentik.stages.prompt.stage import PromptStageView

        return PromptStageView

    @property
    def form(self) -> Type[ModelForm]:
        from authentik.stages.prompt.forms import PromptStageForm

        return PromptStageForm

    def __str__(self):
        return f"Prompt Stage {self.name}"

    class Meta:

        verbose_name = _("Prompt Stage")
        verbose_name_plural = _("Prompt Stages")
