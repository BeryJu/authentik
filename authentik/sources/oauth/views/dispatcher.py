"""Dispatch OAuth views to respective views"""
from django.http import Http404
from django.shortcuts import get_object_or_404
from django.views import View
from structlog import get_logger

from authentik.sources.oauth.models import OAuthSource
from authentik.sources.oauth.types.manager import MANAGER, RequestKind

LOGGER = get_logger()


class DispatcherView(View):
    """Dispatch OAuth Redirect/Callback views to their proper class based on URL parameters"""

    kind = ""

    def dispatch(self, *args, **kwargs):
        """Find Source by slug and forward request"""
        slug = kwargs.get("source_slug", None)
        if not slug:
            raise Http404
        source = get_object_or_404(OAuthSource, slug=slug)
        view = MANAGER.find(source, kind=RequestKind(self.kind))
        LOGGER.debug("dispatching OAuth2 request to", view=view, kind=self.kind)
        return view.as_view()(*args, **kwargs)
