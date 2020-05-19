# Generated by Django 3.0.6 on 2020-05-19 22:07

from django.conf import settings
import django.contrib.auth.models
import django.contrib.auth.validators
import django.contrib.postgres.fields.jsonb
from django.db import migrations, models
import django.db.models.deletion
import django.utils.timezone
import guardian.mixins
import passbook.core.models
import uuid


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("passbook_policies", "0001_initial"),
        ("auth", "0011_update_proxy_permissions"),
    ]

    operations = [
        migrations.CreateModel(
            name="User",
            fields=[
                (
                    "id",
                    models.AutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                ("password", models.CharField(max_length=128, verbose_name="password")),
                (
                    "last_login",
                    models.DateTimeField(
                        blank=True, null=True, verbose_name="last login"
                    ),
                ),
                (
                    "is_superuser",
                    models.BooleanField(
                        default=False,
                        help_text="Designates that this user has all permissions without explicitly assigning them.",
                        verbose_name="superuser status",
                    ),
                ),
                (
                    "username",
                    models.CharField(
                        error_messages={
                            "unique": "A user with that username already exists."
                        },
                        help_text="Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.",
                        max_length=150,
                        unique=True,
                        validators=[
                            django.contrib.auth.validators.UnicodeUsernameValidator()
                        ],
                        verbose_name="username",
                    ),
                ),
                (
                    "first_name",
                    models.CharField(
                        blank=True, max_length=30, verbose_name="first name"
                    ),
                ),
                (
                    "last_name",
                    models.CharField(
                        blank=True, max_length=150, verbose_name="last name"
                    ),
                ),
                (
                    "email",
                    models.EmailField(
                        blank=True, max_length=254, verbose_name="email address"
                    ),
                ),
                (
                    "is_staff",
                    models.BooleanField(
                        default=False,
                        help_text="Designates whether the user can log into this admin site.",
                        verbose_name="staff status",
                    ),
                ),
                (
                    "is_active",
                    models.BooleanField(
                        default=True,
                        help_text="Designates whether this user should be treated as active. Unselect this instead of deleting accounts.",
                        verbose_name="active",
                    ),
                ),
                (
                    "date_joined",
                    models.DateTimeField(
                        default=django.utils.timezone.now, verbose_name="date joined"
                    ),
                ),
                ("uuid", models.UUIDField(default=uuid.uuid4, editable=False)),
                ("name", models.TextField(help_text="User's display name.")),
                ("password_change_date", models.DateTimeField(auto_now_add=True)),
                (
                    "attributes",
                    django.contrib.postgres.fields.jsonb.JSONField(
                        blank=True, default=dict
                    ),
                ),
            ],
            options={"permissions": (("reset_user_password", "Reset Password"),),},
            bases=(guardian.mixins.GuardianUserMixin, models.Model),
            managers=[("objects", django.contrib.auth.models.UserManager()),],
        ),
        migrations.CreateModel(
            name="PropertyMapping",
            fields=[
                (
                    "pm_uuid",
                    models.UUIDField(
                        default=uuid.uuid4,
                        editable=False,
                        primary_key=True,
                        serialize=False,
                    ),
                ),
                ("name", models.TextField()),
                ("expression", models.TextField()),
            ],
            options={
                "verbose_name": "Property Mapping",
                "verbose_name_plural": "Property Mappings",
            },
        ),
        migrations.CreateModel(
            name="Source",
            fields=[
                (
                    "policybindingmodel_ptr",
                    models.OneToOneField(
                        auto_created=True,
                        on_delete=django.db.models.deletion.CASCADE,
                        parent_link=True,
                        primary_key=True,
                        serialize=False,
                        to="passbook_policies.PolicyBindingModel",
                    ),
                ),
                ("name", models.TextField(help_text="Source's display Name.")),
                (
                    "slug",
                    models.SlugField(help_text="Internal source name, used in URLs."),
                ),
                ("enabled", models.BooleanField(default=True)),
                (
                    "property_mappings",
                    models.ManyToManyField(
                        blank=True, default=None, to="passbook_core.PropertyMapping"
                    ),
                ),
            ],
            bases=("passbook_policies.policybindingmodel",),
        ),
        migrations.CreateModel(
            name="UserSourceConnection",
            fields=[
                (
                    "id",
                    models.AutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                ("created", models.DateTimeField(auto_now_add=True)),
                ("last_updated", models.DateTimeField(auto_now=True)),
                (
                    "source",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        to="passbook_core.Source",
                    ),
                ),
                (
                    "user",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        to=settings.AUTH_USER_MODEL,
                    ),
                ),
            ],
            options={"unique_together": {("user", "source")},},
        ),
        migrations.CreateModel(
            name="Token",
            fields=[
                (
                    "token_uuid",
                    models.UUIDField(
                        default=uuid.uuid4,
                        editable=False,
                        primary_key=True,
                        serialize=False,
                    ),
                ),
                (
                    "expires",
                    models.DateTimeField(
                        default=passbook.core.models.default_token_duration
                    ),
                ),
                ("expiring", models.BooleanField(default=True)),
                ("description", models.TextField(blank=True, default="")),
                (
                    "user",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="+",
                        to=settings.AUTH_USER_MODEL,
                    ),
                ),
            ],
            options={"verbose_name": "Token", "verbose_name_plural": "Tokens",},
        ),
        migrations.CreateModel(
            name="Provider",
            fields=[
                (
                    "id",
                    models.AutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                (
                    "property_mappings",
                    models.ManyToManyField(
                        blank=True, default=None, to="passbook_core.PropertyMapping"
                    ),
                ),
            ],
        ),
        migrations.CreateModel(
            name="Group",
            fields=[
                (
                    "group_uuid",
                    models.UUIDField(
                        default=uuid.uuid4,
                        editable=False,
                        primary_key=True,
                        serialize=False,
                    ),
                ),
                ("name", models.CharField(max_length=80, verbose_name="name")),
                (
                    "attributes",
                    django.contrib.postgres.fields.jsonb.JSONField(
                        blank=True, default=dict
                    ),
                ),
                (
                    "parent",
                    models.ForeignKey(
                        blank=True,
                        null=True,
                        on_delete=django.db.models.deletion.SET_NULL,
                        related_name="children",
                        to="passbook_core.Group",
                    ),
                ),
            ],
            options={"unique_together": {("name", "parent")},},
        ),
        migrations.CreateModel(
            name="Application",
            fields=[
                (
                    "policybindingmodel_ptr",
                    models.OneToOneField(
                        auto_created=True,
                        on_delete=django.db.models.deletion.CASCADE,
                        parent_link=True,
                        primary_key=True,
                        serialize=False,
                        to="passbook_policies.PolicyBindingModel",
                    ),
                ),
                ("name", models.TextField(help_text="Application's display Name.")),
                (
                    "slug",
                    models.SlugField(
                        help_text="Internal application name, used in URLs."
                    ),
                ),
                ("skip_authorization", models.BooleanField(default=False)),
                ("meta_launch_url", models.URLField(blank=True, default="")),
                ("meta_icon_url", models.TextField(blank=True, default="")),
                ("meta_description", models.TextField(blank=True, default="")),
                ("meta_publisher", models.TextField(blank=True, default="")),
                (
                    "provider",
                    models.OneToOneField(
                        blank=True,
                        default=None,
                        null=True,
                        on_delete=django.db.models.deletion.SET_DEFAULT,
                        to="passbook_core.Provider",
                    ),
                ),
            ],
            bases=("passbook_policies.policybindingmodel",),
        ),
        migrations.AddField(
            model_name="user",
            name="groups",
            field=models.ManyToManyField(to="passbook_core.Group"),
        ),
        migrations.AddField(
            model_name="user",
            name="sources",
            field=models.ManyToManyField(
                through="passbook_core.UserSourceConnection", to="passbook_core.Source"
            ),
        ),
        migrations.AddField(
            model_name="user",
            name="user_permissions",
            field=models.ManyToManyField(
                blank=True,
                help_text="Specific permissions for this user.",
                related_name="user_set",
                related_query_name="user",
                to="auth.Permission",
                verbose_name="user permissions",
            ),
        ),
    ]
