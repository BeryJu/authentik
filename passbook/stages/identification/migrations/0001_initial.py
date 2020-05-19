# Generated by Django 3.0.6 on 2020-05-19 22:08

import django.contrib.postgres.fields
import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("passbook_flows", "0001_initial"),
    ]

    operations = [
        migrations.CreateModel(
            name="IdentificationStage",
            fields=[
                (
                    "stage_ptr",
                    models.OneToOneField(
                        auto_created=True,
                        on_delete=django.db.models.deletion.CASCADE,
                        parent_link=True,
                        primary_key=True,
                        serialize=False,
                        to="passbook_flows.Stage",
                    ),
                ),
                (
                    "user_fields",
                    django.contrib.postgres.fields.ArrayField(
                        base_field=models.CharField(
                            choices=[("email", "E Mail"), ("username", "Username")],
                            max_length=100,
                        ),
                        help_text="Fields of the user object to match against.",
                        size=None,
                    ),
                ),
                (
                    "template",
                    models.TextField(
                        choices=[
                            ("stages/identification/login.html", "Default Login"),
                            ("stages/identification/recovery.html", "Default Recovery"),
                        ]
                    ),
                ),
            ],
            options={
                "verbose_name": "Identification Stage",
                "verbose_name_plural": "Identification Stages",
            },
            bases=("passbook_flows.stage",),
        ),
    ]
