# Generated by Django 3.0.7 on 2020-06-29 08:57

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("authentik_flows", "0003_auto_20200523_1133"),
    ]

    operations = [
        migrations.AlterField(
            model_name="flow",
            name="designation",
            field=models.CharField(
                choices=[
                    ("authentication", "Authentication"),
                    ("authorization", "Authorization"),
                    ("invalidation", "Invalidation"),
                    ("enrollment", "Enrollment"),
                    ("unenrollment", "Unrenollment"),
                    ("recovery", "Recovery"),
                    ("stage_setup", "Stage Setup"),
                ],
                max_length=100,
            ),
        ),
    ]