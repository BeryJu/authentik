# Generated by Django 3.0.6 on 2020-05-19 22:08

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("passbook_flows", "0001_initial"),
    ]

    operations = [
        migrations.CreateModel(
            name="UserWriteStage",
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
            ],
            options={
                "verbose_name": "User Write Stage",
                "verbose_name_plural": "User Write Stages",
            },
            bases=("passbook_flows.stage",),
        ),
    ]
