"""Metrics view"""
from base64 import b64encode

from django.conf import settings
from django.db import connections
from django.db.utils import OperationalError
from django.http import HttpRequest, HttpResponse
from django.views import View
from django_prometheus.exports import ExportToDjangoView
from django_redis import get_redis_connection
from redis.exceptions import RedisError


class MetricsView(View):
    """Wrapper around ExportToDjangoView, using http-basic auth"""

    def get(self, request: HttpRequest) -> HttpResponse:
        """Check for HTTP-Basic auth"""
        auth_header = request.META.get("HTTP_AUTHORIZATION", "")
        auth_type, _, given_credentials = auth_header.partition(" ")
        credentials = f"monitor:{settings.SECRET_KEY}"
        expected = b64encode(str.encode(credentials)).decode()

        if auth_type != "Basic" or given_credentials != expected:
            response = HttpResponse(status=401)
            response["WWW-Authenticate"] = 'Basic realm="authentik-monitoring"'
            return response

        return ExportToDjangoView(request)


class LiveView(View):
    """View for liveness probe, always returns Http 201"""

    def dispatch(self, request: HttpRequest) -> HttpResponse:
        return HttpResponse(status=201)


class ReadyView(View):
    """View for readiness probe, always returns Http 201, unless sql or redis is down"""

    def dispatch(self, request: HttpRequest) -> HttpResponse:
        try:
            db_conn = connections["default"]
            _ = db_conn.cursor()
        except OperationalError:
            return HttpResponse(status=503)
        try:
            redis_conn = get_redis_connection()
            redis_conn.ping()
        except RedisError:
            return HttpResponse(status=503)
        return HttpResponse(status=201)
