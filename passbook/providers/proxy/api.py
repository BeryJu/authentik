"""ProxyProvider API Views"""
from drf_yasg.utils import swagger_serializer_method
from rest_framework.fields import CharField, ListField, SerializerMethodField
from rest_framework.serializers import ModelSerializer, Serializer
from rest_framework.viewsets import ModelViewSet
from rest_framework.request import Request
from rest_framework.response import Response

from passbook.providers.oauth2.views.provider import ProviderInfoView
from passbook.providers.proxy.models import ProxyProvider


class OpenIDConnectConfigurationSerializer(Serializer):
    """rest_framework Serializer for OIDC Configuration"""

    issuer = CharField()
    authorization_endpoint = CharField()
    token_endpoint = CharField()
    userinfo_endpoint = CharField()
    end_session_endpoint = CharField()
    introspection_endpoint = CharField()
    jwks_uri = CharField()

    response_types_supported = ListField(child=CharField())
    id_token_signing_alg_values_supported = ListField(child=CharField())
    subject_types_supported = ListField(child=CharField())
    token_endpoint_auth_methods_supported = ListField(child=CharField())

    def create(self, request: Request) -> Response:
        raise NotImplementedError

    def update(self, request: Request) -> Response:
        raise NotImplementedError


class ProxyProviderSerializer(ModelSerializer):
    """ProxyProvider Serializer"""

    oidc_configuration = SerializerMethodField()

    def create(self, validated_data):
        instance: ProxyProvider = super().create(validated_data)
        instance.set_oauth_defaults()
        instance.save()
        return instance

    def update(self, instance: ProxyProvider, validated_data):
        instance.set_oauth_defaults()
        return super().update(instance, validated_data)

    class Meta:

        model = ProxyProvider
        fields = [
            "pk",
            "name",
            "internal_host",
            "external_host",
            "client_id",
            "client_secret",
            "oidc_configuration",
            "cookie_secret",
            "certificate",
        ]

    @swagger_serializer_method(serializer_or_field=OpenIDConnectConfigurationSerializer)
    def get_oidc_configuration(self, obj: ProxyProvider):
        """Embed OpenID Connect provider information"""
        # pylint: disable=protected-access
        return ProviderInfoView(request=self.context["request"]._request).get_info(obj)


class ProxyProviderViewSet(ModelViewSet):
    """ProxyProvider Viewset"""

    queryset = ProxyProvider.objects.all()
    serializer_class = ProxyProviderSerializer
