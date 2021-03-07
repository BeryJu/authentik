"""API Utilities"""
from django.db.models import Model
from rest_framework.fields import CharField, IntegerField
from rest_framework.serializers import Serializer, SerializerMethodField


class MetaNameSerializer(Serializer):
    """Add verbose names to response"""

    verbose_name = SerializerMethodField()
    verbose_name_plural = SerializerMethodField()

    def create(self, validated_data: dict) -> Model:
        raise NotImplementedError

    def update(self, instance: Model, validated_data: dict) -> Model:
        raise NotImplementedError

    def get_verbose_name(self, obj: Model) -> str:
        """Return object's verbose_name"""
        return obj._meta.verbose_name

    def get_verbose_name_plural(self, obj: Model) -> str:
        """Return object's plural verbose_name"""
        return obj._meta.verbose_name_plural


class TypeCreateSerializer(Serializer):
    """Types of an object that can be created"""

    name = CharField(required=True)
    description = CharField(required=True)
    link = CharField(required=True)

    def create(self, validated_data: dict) -> Model:
        raise NotImplementedError

    def update(self, instance: Model, validated_data: dict) -> Model:
        raise NotImplementedError


class CacheSerializer(Serializer):
    """Generic cache stats for an object"""

    count = IntegerField(read_only=True)

    def create(self, validated_data: dict) -> Model:
        raise NotImplementedError

    def update(self, instance: Model, validated_data: dict) -> Model:
        raise NotImplementedError
