"""Proxy Provider Docker Contoller"""
from typing import Dict
from urllib.parse import urlparse

from authentik.outposts.controllers.docker import DockerController
from authentik.outposts.models import DockerServiceConnection, Outpost
from authentik.providers.proxy.models import ProxyProvider


class ProxyDockerController(DockerController):
    """Proxy Provider Docker Contoller"""

    def __init__(self, outpost: Outpost, connection: DockerServiceConnection):
        super().__init__(outpost, connection)
        self.deployment_ports = {
            "http": 4180,
            "https": 4443,
        }

    def _get_labels(self) -> Dict[str, str]:
        hosts = []
        for proxy_provider in ProxyProvider.objects.filter(outpost__in=[self.outpost]):
            proxy_provider: ProxyProvider
            external_host_name = urlparse(proxy_provider.external_host)
            hosts.append(f"`{external_host_name}`")
        traefik_name = f"pb-outpost-{self.outpost.pk.hex}"
        return {
            "traefik.enable": "true",
            f"traefik.http.routers.{traefik_name}-router.rule": f"Host({','.join(hosts)})",
            f"traefik.http.routers.{traefik_name}-router.tls": "true",
            f"traefik.http.routers.{traefik_name}-router.service": f"{traefik_name}-service",
            f"traefik.http.services.{traefik_name}-service.loadbalancer.healthcheck.path": "/",
            f"traefik.http.services.{traefik_name}-service.loadbalancer.server.port": "4180",
        }
