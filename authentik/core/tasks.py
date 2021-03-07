"""authentik core tasks"""
from datetime import datetime
from io import StringIO

from boto3.exceptions import Boto3Error
from botocore.exceptions import BotoCoreError, ClientError
from dbbackup.db.exceptions import CommandConnectorError
from django.contrib.humanize.templatetags.humanize import naturaltime
from django.core import management
from django.utils.timezone import now
from structlog.stdlib import get_logger

from authentik.core.models import ExpiringModel
from authentik.events.monitored_tasks import MonitoredTask, TaskResult, TaskResultStatus
from authentik.root.celery import CELERY_APP

LOGGER = get_logger()


@CELERY_APP.task(bind=True, base=MonitoredTask)
def clean_expired_models(self: MonitoredTask):
    """Remove expired objects"""
    messages = []
    for cls in ExpiringModel.__subclasses__():
        cls: ExpiringModel
        amount, _ = (
            cls.objects.all()
            .exclude(expiring=False)
            .exclude(expiring=True, expires__gt=now())
            .delete()
        )
        LOGGER.debug("Deleted expired models", model=cls, amount=amount)
        messages.append(f"Deleted {amount} expired {cls._meta.verbose_name_plural}")
    self.set_status(TaskResult(TaskResultStatus.SUCCESSFUL, messages))


@CELERY_APP.task(bind=True, base=MonitoredTask)
def backup_database(self: MonitoredTask):  # pragma: no cover
    """Database backup"""
    self.result_timeout_hours = 25
    try:
        start = datetime.now()
        out = StringIO()
        management.call_command("dbbackup", quiet=True, stdout=out)
        self.set_status(
            TaskResult(
                TaskResultStatus.SUCCESSFUL,
                [
                    f"Successfully finished database backup {naturaltime(start)} {out.getvalue()}",
                ],
            )
        )
        LOGGER.info("Successfully backed up database.")
    except (
        IOError,
        BotoCoreError,
        ClientError,
        Boto3Error,
        PermissionError,
        CommandConnectorError,
    ) as exc:
        self.set_status(TaskResult(TaskResultStatus.ERROR).with_error(exc))
