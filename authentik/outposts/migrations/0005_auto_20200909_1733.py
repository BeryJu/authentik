# Generated by Django 3.1.1 on 2020-09-09 17:33

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("authentik_outposts", "0004_auto_20200830_1056"),
    ]

    operations = [
        migrations.AlterField(
            model_name="outpost",
            name="deployment_type",
            field=models.TextField(
                choices=[("custom", "Custom")],
                default="custom",
                help_text="Select between authentik-managed deployment types or a custom deployment.",
            ),
        ),
    ]
