# Generated by Django 3.1.6 on 2021-02-13 16:40

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("authentik_policies_event_matcher", "0007_auto_20210209_1657"),
    ]

    operations = [
        migrations.AlterField(
            model_name="eventmatcherpolicy",
            name="app",
            field=models.TextField(
                blank=True,
                choices=[
                    ("authentik.admin", "authentik Admin"),
                    ("authentik.api", "authentik API"),
                    ("authentik.events", "authentik Events"),
                    ("authentik.crypto", "authentik Crypto"),
                    ("authentik.flows", "authentik Flows"),
                    ("authentik.outposts", "authentik Outpost"),
                    ("authentik.lib", "authentik lib"),
                    ("authentik.policies", "authentik Policies"),
                    ("authentik.policies.dummy", "authentik Policies.Dummy"),
                    (
                        "authentik.policies.event_matcher",
                        "authentik Policies.Event Matcher",
                    ),
                    ("authentik.policies.expiry", "authentik Policies.Expiry"),
                    ("authentik.policies.expression", "authentik Policies.Expression"),
                    (
                        "authentik.policies.group_membership",
                        "authentik Policies.Group Membership",
                    ),
                    ("authentik.policies.hibp", "authentik Policies.HaveIBeenPwned"),
                    ("authentik.policies.password", "authentik Policies.Password"),
                    ("authentik.policies.reputation", "authentik Policies.Reputation"),
                    ("authentik.providers.proxy", "authentik Providers.Proxy"),
                    ("authentik.providers.oauth2", "authentik Providers.OAuth2"),
                    ("authentik.providers.saml", "authentik Providers.SAML"),
                    ("authentik.recovery", "authentik Recovery"),
                    ("authentik.sources.ldap", "authentik Sources.LDAP"),
                    ("authentik.sources.oauth", "authentik Sources.OAuth"),
                    ("authentik.sources.saml", "authentik Sources.SAML"),
                    ("authentik.stages.captcha", "authentik Stages.Captcha"),
                    ("authentik.stages.consent", "authentik Stages.Consent"),
                    ("authentik.stages.dummy", "authentik Stages.Dummy"),
                    ("authentik.stages.email", "authentik Stages.Email"),
                    ("authentik.stages.prompt", "authentik Stages.Prompt"),
                    (
                        "authentik.stages.identification",
                        "authentik Stages.Identification",
                    ),
                    ("authentik.stages.invitation", "authentik Stages.User Invitation"),
                    ("authentik.stages.user_delete", "authentik Stages.User Delete"),
                    ("authentik.stages.user_login", "authentik Stages.User Login"),
                    ("authentik.stages.user_logout", "authentik Stages.User Logout"),
                    ("authentik.stages.user_write", "authentik Stages.User Write"),
                    ("authentik.stages.otp_static", "authentik Stages.OTP.Static"),
                    ("authentik.stages.otp_time", "authentik Stages.OTP.Time"),
                    ("authentik.stages.otp_validate", "authentik Stages.OTP.Validate"),
                    ("authentik.stages.password", "authentik Stages.Password"),
                    ("authentik.stages.webauthn", "authentik Stages.WebAuthN"),
                    ("authentik.managed", "authentik Managed"),
                    ("authentik.core", "authentik Core"),
                ],
                default="",
                help_text="Match events created by selected application. When left empty, all applications are matched.",
            ),
        ),
    ]
