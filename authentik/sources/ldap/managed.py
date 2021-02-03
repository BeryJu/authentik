"""LDAP Source managed objects"""
from authentik.managed.manager import EnsureExists, ObjectManager
from authentik.sources.ldap.models import LDAPPropertyMapping


class LDAPProviderManager(ObjectManager):
    """LDAP Source managed objects"""

    def reconcile(self):
        return [
            EnsureExists(
                LDAPPropertyMapping,
                "object_field",
                name="authentik default LDAP Mapping: Name",
                object_field="name",
                expression="return ldap.get('name')",
            ),
            EnsureExists(
                LDAPPropertyMapping,
                "object_field",
                name="authentik default LDAP Mapping: mail",
                object_field="email",
                expression="return ldap.get('mail')",
            ),
            EnsureExists(
                LDAPPropertyMapping,
                "object_field",
                name="authentik default Active Directory Mapping: sAMAccountName",
                object_field="username",
                expression="return ldap.get('sAMAccountName')",
            ),
            EnsureExists(
                LDAPPropertyMapping,
                "object_field",
                name="authentik default Active Directory Mapping: userPrincipalName",
                object_field="attributes.upn",
                expression="return ldap.get('userPrincipalName')",
            ),
        ]
