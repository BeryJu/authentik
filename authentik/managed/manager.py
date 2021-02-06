"""Managed objects manager"""
from typing import Type

from structlog.stdlib import get_logger

from authentik.managed.models import ManagedModel

LOGGER = get_logger()


class EnsureOp:
    """Ensure operation, executed as part of an ObjectManager run"""

    _obj: Type[ManagedModel]
    _managed_uid: str
    _kwargs: dict

    def __init__(self, obj: Type[ManagedModel], managed_uid: str, **kwargs) -> None:
        self._obj = obj
        self._managed_uid = managed_uid
        self._kwargs = kwargs

    def run(self):
        """Do the actual ensure action"""
        raise NotImplementedError


class EnsureExists(EnsureOp):
    """Ensure object exists, with kwargs as given values"""

    def run(self):
        self._kwargs.setdefault("managed", self._managed_uid)
        self._obj.objects.update_or_create(
            **{
                "managed": self._managed_uid,
                "defaults": self._kwargs,
            }
        )


class ObjectManager:
    """Base class for Apps Object manager"""

    def run(self):
        """Main entrypoint for tasks, iterate through all implementation of this
        and execute all operations"""
        for sub in ObjectManager.__subclasses__():
            sub_inst = sub()
            ops = sub_inst.reconcile()
            LOGGER.debug("Reconciling managed objects", manager=sub.__name__)
            for operation in ops:
                operation.run()

    def reconcile(self) -> list[EnsureOp]:
        """Method which is implemented in subclass that returns a list of Operations"""
        raise NotImplementedError
