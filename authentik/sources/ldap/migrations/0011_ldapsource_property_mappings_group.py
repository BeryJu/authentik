# Generated by Django 3.1.6 on 2021-02-06 14:01

from django.apps.registry import Apps
from django.db import migrations, models


def set_default_group_mappings(apps: Apps, schema_editor):
    LDAPPropertyMapping = apps.get_model(
        "authentik_sources_ldap", "LDAPPropertyMapping"
    )
    LDAPSource = apps.get_model("authentik_sources_ldap", "LDAPSource")
    db_alias = schema_editor.connection.alias

    for source in LDAPSource.objects.using(db_alias).all():
        if source.property_mappings_group.exists():
            continue
        source.property_mappings_group.set(
            LDAPPropertyMapping.objects.using(db_alias).filter(
                managed="goauthentik.io/sources/ldap/default-name"
            )
        )
        source.save()


class Migration(migrations.Migration):

    dependencies = [
        ("authentik_sources_ldap", "0010_auto_20210205_1027"),
    ]

    operations = [
        migrations.AddField(
            model_name="ldapsource",
            name="property_mappings_group",
            field=models.ManyToManyField(
                blank=True,
                default=None,
                help_text="Property mappings used for group creation/updating.",
                to="authentik_core.PropertyMapping",
            ),
        ),
        migrations.RunPython(set_default_group_mappings),
    ]