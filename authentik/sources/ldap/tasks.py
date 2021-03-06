"""LDAP Sync tasks"""
from django.utils.text import slugify
from ldap3.core.exceptions import LDAPException
from structlog.stdlib import get_logger

from authentik.events.monitored_tasks import MonitoredTask, TaskResult, TaskResultStatus
from authentik.root.celery import CELERY_APP
from authentik.sources.ldap.models import LDAPSource
from authentik.sources.ldap.sync.groups import GroupLDAPSynchronizer
from authentik.sources.ldap.sync.membership import MembershipLDAPSynchronizer
from authentik.sources.ldap.sync.users import UserLDAPSynchronizer

LOGGER = get_logger()


@CELERY_APP.task()
def ldap_sync_all():
    """Sync all sources"""
    for source in LDAPSource.objects.filter(enabled=True):
        ldap_sync.delay(source.pk)


@CELERY_APP.task(bind=True, base=MonitoredTask)
def ldap_sync(self: MonitoredTask, source_pk: str):
    """Synchronization of an LDAP Source"""
    self.result_timeout_hours = 2
    try:
        source: LDAPSource = LDAPSource.objects.get(pk=source_pk)
    except LDAPSource.DoesNotExist:
        # Because the source couldn't be found, we don't have a UID
        # to set the state with
        return
    self.set_uid(slugify(source.name))
    try:
        messages = []
        for sync_class in [
            UserLDAPSynchronizer,
            GroupLDAPSynchronizer,
            MembershipLDAPSynchronizer,
        ]:
            sync_inst = sync_class(source)
            count = sync_inst.sync()
            messages.append(f"Synced {count} objects from {sync_class.__name__}")
        self.set_status(
            TaskResult(
                TaskResultStatus.SUCCESSFUL,
                messages,
            )
        )
    except LDAPException as exc:
        # No explicit event is created here as .set_status with an error will do that
        LOGGER.debug(exc)
        self.set_status(TaskResult(TaskResultStatus.ERROR).with_error(exc))
