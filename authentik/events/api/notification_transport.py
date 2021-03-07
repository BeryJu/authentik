"""NotificationTransport API Views"""
from django.http.response import Http404
from drf_yasg2.utils import no_body, swagger_auto_schema
from guardian.shortcuts import get_objects_for_user
from rest_framework.decorators import action
from rest_framework.fields import CharField, ListField, SerializerMethodField
from rest_framework.request import Request
from rest_framework.response import Response
from rest_framework.serializers import ModelSerializer, Serializer
from rest_framework.viewsets import ModelViewSet

from authentik.events.models import (
    Notification,
    NotificationSeverity,
    NotificationTransport,
    NotificationTransportError,
    TransportMode,
)


class NotificationTransportSerializer(ModelSerializer):
    """NotificationTransport Serializer"""

    mode_verbose = SerializerMethodField()

    def get_mode_verbose(self, instance: NotificationTransport):
        """Return selected mode with a UI Label"""
        return TransportMode(instance.mode).label

    class Meta:

        model = NotificationTransport
        fields = [
            "pk",
            "name",
            "mode",
            "mode_verbose",
            "webhook_url",
        ]


class NotificationTransportTestSerializer(Serializer):

    messages = ListField(child=CharField())


class NotificationTransportViewSet(ModelViewSet):
    """NotificationTransport Viewset"""

    queryset = NotificationTransport.objects.all()
    serializer_class = NotificationTransportSerializer

    @swagger_auto_schema(
        responses={200: NotificationTransportTestSerializer(many=False)},
        request_body=no_body,
    )
    @action(detail=True, methods=["post"])
    # pylint: disable=invalid-name
    def test(self, request: Request, pk=None) -> Response:
        """Send example notification using selected transport. Requires
        Modify permissions."""
        transports = get_objects_for_user(
            request.user, "authentik_events.change_notificationtransport"
        ).filter(pk=pk)
        if not transports.exists():
            raise Http404
        transport: NotificationTransport = transports.first()
        notification = Notification(
            severity=NotificationSeverity.NOTICE,
            body=f"Test Notification from transport {transport.name}",
            user=request.user,
        )
        try:
            response = NotificationTransportTestSerializer(
                data={"messages": transport.send(notification)}
            )
            response.is_valid()
            return Response(response.data)
        except NotificationTransportError as exc:
            return Response(str(exc.__cause__ or None), status=503)
