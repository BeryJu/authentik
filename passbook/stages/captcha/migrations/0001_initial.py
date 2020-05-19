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
            name="CaptchaStage",
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
                    "public_key",
                    models.TextField(
                        help_text="Public key, acquired from https://www.google.com/recaptcha/intro/v3.html"
                    ),
                ),
                (
                    "private_key",
                    models.TextField(
                        help_text="Private key, acquired from https://www.google.com/recaptcha/intro/v3.html"
                    ),
                ),
            ],
            options={
                "verbose_name": "Captcha Stage",
                "verbose_name_plural": "Captcha Stages",
            },
            bases=("passbook_flows.stage",),
        ),
    ]
