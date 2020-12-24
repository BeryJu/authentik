"""authentik events models"""

from inspect import getmodule, stack
from typing import Optional, Union
from uuid import uuid4

from django.conf import settings
from django.core.exceptions import ValidationError
from django.db import models
from django.http import HttpRequest
from django.utils.translation import gettext as _
from structlog import get_logger

from authentik.core.middleware import (
    SESSION_IMPERSONATE_ORIGINAL_USER,
    SESSION_IMPERSONATE_USER,
)
from authentik.core.models import User
from authentik.events.utils import cleanse_dict, get_user, sanitize_dict
from authentik.lib.utils.http import get_client_ip
from authentik.policies.models import PolicyBindingModel

LOGGER = get_logger("authentik.events")


class EventAction(models.TextChoices):
    """All possible actions to save into the events log"""

    LOGIN = "login"
    LOGIN_FAILED = "login_failed"
    LOGOUT = "logout"

    USER_WRITE = "user_write"
    SUSPICIOUS_REQUEST = "suspicious_request"
    PASSWORD_SET = "password_set"  # noqa # nosec

    TOKEN_VIEW = "token_view"  # nosec

    INVITE_CREATED = "invitation_created"
    INVITE_USED = "invitation_used"

    AUTHORIZE_APPLICATION = "authorize_application"
    SOURCE_LINKED = "source_linked"

    IMPERSONATION_STARTED = "impersonation_started"
    IMPERSONATION_ENDED = "impersonation_ended"

    POLICY_EXECUTION = "policy_execution"
    POLICY_EXCEPTION = "policy_exception"
    PROPERTY_MAPPING_EXCEPTION = "property_mapping_exception"

    MODEL_CREATED = "model_created"
    MODEL_UPDATED = "model_updated"
    MODEL_DELETED = "model_deleted"

    UPDATE_AVAILABLE = "update_available"

    CUSTOM_PREFIX = "custom_"


class Event(models.Model):
    """An individual Audit/Metrics/Notification/Error Event"""

    event_uuid = models.UUIDField(primary_key=True, editable=False, default=uuid4)
    user = models.JSONField(default=dict)
    action = models.TextField(choices=EventAction.choices)
    app = models.TextField()
    context = models.JSONField(default=dict, blank=True)
    client_ip = models.GenericIPAddressField(null=True)
    created = models.DateTimeField(auto_now_add=True)

    @staticmethod
    def _get_app_from_request(request: HttpRequest) -> str:
        if not isinstance(request, HttpRequest):
            return ""
        return request.resolver_match.app_name

    @staticmethod
    def new(
        action: Union[str, EventAction],
        app: Optional[str] = None,
        _inspect_offset: int = 1,
        **kwargs,
    ) -> "Event":
        """Create new Event instance from arguments. Instance is NOT saved."""
        if not isinstance(action, EventAction):
            action = EventAction.CUSTOM_PREFIX + action
        if not app:
            app = getmodule(stack()[_inspect_offset][0]).__name__
        cleaned_kwargs = cleanse_dict(sanitize_dict(kwargs))
        event = Event(action=action, app=app, context=cleaned_kwargs)
        return event

    def set_user(self, user: User) -> "Event":
        """Set `.user` based on user, ensuring the correct attributes are copied.
        This should only be used when self.from_http is *not* used."""
        self.user = get_user(user)
        return self

    def from_http(
        self, request: HttpRequest, user: Optional[settings.AUTH_USER_MODEL] = None
    ) -> "Event":
        """Add data from a Django-HttpRequest, allowing the creation of
        Events independently from requests.
        `user` arguments optionally overrides user from requests."""
        if hasattr(request, "user"):
            original_user = None
            if hasattr(request, "session"):
                original_user = request.session.get(
                    SESSION_IMPERSONATE_ORIGINAL_USER, None
                )
            self.user = get_user(request.user, original_user)
        if user:
            self.user = get_user(user)
        # Check if we're currently impersonating, and add that user
        if hasattr(request, "session"):
            if SESSION_IMPERSONATE_ORIGINAL_USER in request.session:
                self.user = get_user(request.session[SESSION_IMPERSONATE_ORIGINAL_USER])
                self.user["on_behalf_of"] = get_user(
                    request.session[SESSION_IMPERSONATE_USER]
                )
        # User 255.255.255.255 as fallback if IP cannot be determined
        self.client_ip = get_client_ip(request) or "255.255.255.255"
        # If there's no app set, we get it from the requests too
        if not self.app:
            self.app = Event._get_app_from_request(request)
        self.save()
        return self

    def save(self, *args, **kwargs):
        if not self._state.adding:
            raise ValidationError(
                "you may not edit an existing %s" % self._meta.model_name
            )
        LOGGER.debug(
            "Created Event",
            action=self.action,
            context=self.context,
            client_ip=self.client_ip,
            user=self.user,
        )
        return super().save(*args, **kwargs)

    def __str__(self) -> str:
        return f"<Event action={self.action} user={self.user} context={self.context}>"

    class Meta:

        verbose_name = _("Event")
        verbose_name_plural = _("Events")


class EventAlertAction(models.Model):
    """Action which is executed when a Trigger matches"""

    name = models.TextField(unique=True)

    def execute(self, event: Event):
        """execute which is executed when alert trigger matches"""
        # TODO: do execute


class EventAlertTrigger(PolicyBindingModel):
    """Alert which is triggered when a certain criteria matches against Events"""

    name = models.TextField(unique=True)
    action = models.ForeignKey(
        EventAlertAction,
        on_delete=models.SET_DEFAULT,
        default=None,
        blank=True,
        null=True,
    )

    class Meta:

        verbose_name = _("Event Alert Trigger")
        verbose_name_plural = _("Event Alert Triggers")