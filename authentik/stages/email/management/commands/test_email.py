"""Send a test-email with global settings"""
from uuid import uuid4

from django.core.management.base import BaseCommand, no_translations

from authentik.stages.email.models import EmailStage
from authentik.stages.email.tasks import send_mail
from authentik.stages.email.utils import TemplateEmailMessage


class Command(BaseCommand):  # pragma: no cover
    """Send a test-email with global settings"""

    @no_translations
    def handle(self, *args, **options):
        """Send a test-email with global settings"""
        delete_stage = False
        if options["stage"]:
            stage = EmailStage.objects.get(name=options["stage"])
        else:
            stage = EmailStage.objects.create(
                name=f"temp-global-stage-{uuid4()}", use_global_settings=True
            )
            delete_stage = True
        message = TemplateEmailMessage(
            subject="authentik Test-Email",
            template_name="email/setup.html",
            to=[options["to"]],
            template_context={},
        )
        try:
            # pyright: reportGeneralTypeIssues=false
            send_mail(  # pylint: disable=no-value-for-parameter
                stage.pk, message.__dict__
            )
        finally:
            if delete_stage:
                stage.delete()

    def add_arguments(self, parser):
        parser.add_argument("to", type=str)
        parser.add_argument("-s", "--stage", type=str)
