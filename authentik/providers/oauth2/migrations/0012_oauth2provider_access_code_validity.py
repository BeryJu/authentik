# Generated by Django 3.2 on 2021-04-28 18:17

from django.db import migrations, models

import authentik.lib.utils.time


class Migration(migrations.Migration):

    dependencies = [
        ("authentik_providers_oauth2", "0011_managed"),
    ]

    operations = [
        migrations.AddField(
            model_name="oauth2provider",
            name="access_code_validity",
            field=models.TextField(
                default="minutes=1",
                help_text="Access codes not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).",
                validators=[authentik.lib.utils.time.timedelta_string_validator],
            ),
        ),
    ]
