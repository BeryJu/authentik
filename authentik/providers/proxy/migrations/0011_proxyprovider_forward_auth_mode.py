# Generated by Django 3.2 on 2021-04-27 18:47

from django.db import migrations, models

import authentik.lib.models


class Migration(migrations.Migration):

    dependencies = [
        ("authentik_providers_proxy", "0010_auto_20201214_0942"),
    ]

    operations = [
        migrations.AddField(
            model_name="proxyprovider",
            name="forward_auth_mode",
            field=models.BooleanField(
                default=False,
                help_text="Enable support for forwardAuth in traefik and nginx auth_request. Exclusive with internal_host.",
            ),
        ),
        migrations.AlterField(
            model_name="proxyprovider",
            name="internal_host",
            field=models.TextField(
                blank=True,
                validators=[
                    authentik.lib.models.DomainlessURLValidator(
                        schemes=("http", "https")
                    )
                ],
            ),
        ),
    ]
