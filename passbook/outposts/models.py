"""Outpost models"""
from dataclasses import asdict, dataclass
from datetime import datetime
from typing import Iterable, Optional
from uuid import uuid4

from django.contrib.postgres.fields import ArrayField
from django.core.cache import cache
from django.db import models
from guardian.shortcuts import assign_perm

from passbook.core.models import Provider, Token, TokenIntents, User
from passbook.lib.config import CONFIG


@dataclass
class OutpostConfig:
    """Configuration an outpost uses to configure it self"""

    log_level: str = CONFIG.y("log_level")
    error_reporting_enabled: bool = CONFIG.y_bool("error_reporting.enabled")
    error_reporting_environment: str = CONFIG.y(
        "error_reporting.environment", "customer"
    )


class OutpostModel:
    """Base model for providers that need more objects than just themselves"""

    def get_required_objects(self) -> Iterable[models.Model]:
        """Return a list of all required objects"""
        return [self]


class OutpostType(models.TextChoices):
    """Outpost types, currently only the reverse proxy is available"""

    PROXY = "proxy"


def default_outpost_config():
    """Get default outpost config"""
    return asdict(OutpostConfig())


class Outpost(models.Model):
    """Outpost instance which manages a service user and token"""

    uuid = models.UUIDField(default=uuid4, editable=False, primary_key=True)

    name = models.TextField()

    type = models.TextField(choices=OutpostType.choices, default=OutpostType.PROXY)

    providers = models.ManyToManyField(Provider)

    channels = ArrayField(models.TextField(), default=list)

    config = models.JSONField(default=default_outpost_config)

    @property
    def health_cache_key(self) -> str:
        """Key by which the outposts health status is saved"""
        return f"outpost_{self.uuid.hex}_health"

    @property
    def health(self) -> Optional[datetime]:
        """Get outpost's health status"""
        key = self.health_cache_key
        value = cache.get(key, None)
        if value:
            return datetime.fromtimestamp(value)
        return None

    def _create_user(self) -> User:
        """Create user and assign permissions for all required objects"""
        user: User = User.objects.create(username=f"pb-outpost-{self.uuid.hex}")
        user.set_unusable_password()
        user.save()
        for model in self.get_required_objects():
            assign_perm(
                f"{model._meta.app_label}.view_{model._meta.model_name}", user, model
            )
        return user

    @property
    def user(self) -> User:
        """Get/create user with access to all required objects"""
        user = User.objects.filter(username=f"pb-outpost-{self.uuid.hex}")
        if user.exists():
            return user.first()
        return self._create_user()

    @property
    def token(self) -> Token:
        """Get/create token for auto-generated user"""
        token = Token.filter_not_expired(user=self.user, intent=TokenIntents.INTENT_API)
        if token.exists():
            return token
        return Token.objects.create(
            user=self.user,
            intent=TokenIntents.INTENT_API,
            description=f"Autogenerated by passbook for Outpost {self.name}",
            expiring=False,
        )

    def get_required_objects(self) -> Iterable[models.Model]:
        """Get an iterator of all objects the user needs read access to"""
        objects = [self]
        for provider in (
            Provider.objects.filter(outpost=self).select_related().select_subclasses()
        ):
            if isinstance(provider, OutpostModel):
                objects.extend(provider.get_required_objects())
            else:
                objects.append(provider)
        return objects

    def __str__(self) -> str:
        return f"Outpost {self.name}"
