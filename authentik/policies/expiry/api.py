"""Password Expiry Policy API Views"""
from rest_framework.serializers import ModelSerializer
from rest_framework.viewsets import ModelViewSet

from authentik.policies.expiry.models import PasswordExpiryPolicy
from authentik.policies.forms import GENERAL_SERIALIZER_FIELDS


class PasswordExpiryPolicySerializer(ModelSerializer):
    """Password Expiry Policy Serializer"""

    class Meta:
        model = PasswordExpiryPolicy
        fields = GENERAL_SERIALIZER_FIELDS + ["days", "deny_only"]


class PasswordExpiryPolicyViewSet(ModelViewSet):
    """Password Expiry Viewset"""

    queryset = PasswordExpiryPolicy.objects.all()
    serializer_class = PasswordExpiryPolicySerializer
