# Generated by Django 3.1.4 on 2020-12-24 10:32

import django.db.models.deletion
from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("authentik_policies", "0004_policy_execution_logging"),
    ]

    operations = [
        migrations.CreateModel(
            name="EventMatcherPolicy",
            fields=[
                (
                    "policy_ptr",
                    models.OneToOneField(
                        auto_created=True,
                        on_delete=django.db.models.deletion.CASCADE,
                        parent_link=True,
                        primary_key=True,
                        serialize=False,
                        to="authentik_policies.policy",
                    ),
                ),
                (
                    "action",
                    models.TextField(
                        blank=True,
                        choices=[
                            ("login", "Login"),
                            ("login_failed", "Login Failed"),
                            ("logout", "Logout"),
                            ("user_write", "User Write"),
                            ("suspicious_request", "Suspicious Request"),
                            ("password_set", "Password Set"),
                            ("token_view", "Token View"),
                            ("invitation_created", "Invite Created"),
                            ("invitation_used", "Invite Used"),
                            ("authorize_application", "Authorize Application"),
                            ("source_linked", "Source Linked"),
                            ("impersonation_started", "Impersonation Started"),
                            ("impersonation_ended", "Impersonation Ended"),
                            ("policy_execution", "Policy Execution"),
                            ("policy_exception", "Policy Exception"),
                            (
                                "property_mapping_exception",
                                "Property Mapping Exception",
                            ),
                            ("model_created", "Model Created"),
                            ("model_updated", "Model Updated"),
                            ("model_deleted", "Model Deleted"),
                            ("update_available", "Update Available"),
                            ("custom_", "Custom Prefix"),
                        ],
                    ),
                ),
                ("client_ip", models.TextField(blank=True)),
            ],
            options={
                "verbose_name": "Group Membership Policy",
                "verbose_name_plural": "Group Membership Policies",
            },
            bases=("authentik_policies.policy",),
        ),
    ]
