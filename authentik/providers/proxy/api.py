"""ProxyProvider API Views"""
from drf_yasg.utils import swagger_serializer_method
from rest_framework.fields import CharField, ListField, SerializerMethodField
from rest_framework.serializers import ModelSerializer
from rest_framework.viewsets import ModelViewSet, ReadOnlyModelViewSet

from authentik.core.api.providers import ProviderSerializer
from authentik.core.api.utils import PassiveSerializer
from authentik.providers.oauth2.views.provider import ProviderInfoView
from authentik.providers.proxy.models import ProxyProvider


class OpenIDConnectConfigurationSerializer(PassiveSerializer):
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


class ProxyProviderSerializer(ProviderSerializer):
    """ProxyProvider Serializer"""

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
        fields = ProviderSerializer.Meta.fields + [
            "internal_host",
            "external_host",
            "internal_host_ssl_validation",
            "certificate",
            "skip_path_regex",
            "basic_auth_enabled",
            "basic_auth_password_attribute",
            "basic_auth_user_attribute",
            "forward_auth_mode",
        ]


class ProxyProviderViewSet(ModelViewSet):
    """ProxyProvider Viewset"""

    queryset = ProxyProvider.objects.all()
    serializer_class = ProxyProviderSerializer
    ordering = ["name"]


class ProxyOutpostConfigSerializer(ModelSerializer):
    """ProxyProvider Serializer"""

    oidc_configuration = SerializerMethodField()

    class Meta:

        model = ProxyProvider
        fields = [
            "pk",
            "name",
            "internal_host",
            "external_host",
            "internal_host_ssl_validation",
            "client_id",
            "client_secret",
            "oidc_configuration",
            "cookie_secret",
            "certificate",
            "skip_path_regex",
            "basic_auth_enabled",
            "basic_auth_password_attribute",
            "basic_auth_user_attribute",
            "forward_auth_mode",
        ]

    @swagger_serializer_method(serializer_or_field=OpenIDConnectConfigurationSerializer)
    def get_oidc_configuration(self, obj: ProxyProvider):
        """Embed OpenID Connect provider information"""
        return ProviderInfoView(request=self.context["request"]._request).get_info(obj)


class ProxyOutpostConfigViewSet(ReadOnlyModelViewSet):
    """ProxyProvider Viewset"""

    queryset = ProxyProvider.objects.filter(application__isnull=False)
    serializer_class = ProxyOutpostConfigSerializer
    ordering = ["name"]
