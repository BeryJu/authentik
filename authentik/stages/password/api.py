"""PasswordStage API Views"""
from rest_framework.serializers import ModelSerializer
from rest_framework.viewsets import ModelViewSet

from authentik.stages.password.models import PasswordStage


class PasswordStageSerializer(ModelSerializer):
    """PasswordStage Serializer"""

    class Meta:

        model = PasswordStage
        fields = [
            "pk",
            "name",
            "backends",
            "configure_flow",
            "failed_attempts_before_cancel",
        ]


class PasswordStageViewSet(ModelViewSet):
    """PasswordStage Viewset"""

    queryset = PasswordStage.objects.all()
    serializer_class = PasswordStageSerializer
