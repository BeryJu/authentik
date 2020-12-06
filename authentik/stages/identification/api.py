"""Identification Stage API Views"""
from rest_framework.serializers import ModelSerializer
from rest_framework.viewsets import ModelViewSet

from authentik.stages.identification.models import IdentificationStage


class IdentificationStageSerializer(ModelSerializer):
    """IdentificationStage Serializer"""

    class Meta:

        model = IdentificationStage
        fields = [
            "pk",
            "name",
            "user_fields",
            "case_insensitive_matching",
            "show_matched_user",
            "template",
            "enrollment_flow",
            "recovery_flow",
        ]


class IdentificationStageViewSet(ModelViewSet):
    """IdentificationStage Viewset"""

    queryset = IdentificationStage.objects.all()
    serializer_class = IdentificationStageSerializer
