# Generated by Django 3.0.6 on 2020-05-23 15:47

from django.apps.registry import Apps
from django.db import migrations
from django.db.backends.base.schema import BaseDatabaseSchemaEditor

from passbook.flows.models import FlowDesignation
from passbook.stages.prompt.models import FieldTypes

FLOW_POLICY_EXPRESSION = """return pb_is_sso_flow"""
PROMPT_POLICY_EXPRESSION = (
    """return 'username' in pb_flow_plan.context['prompt_data']"""
)


def create_default_source_enrollment_flow(
    apps: Apps, schema_editor: BaseDatabaseSchemaEditor
):
    Flow = apps.get_model("passbook_flows", "Flow")
    FlowStageBinding = apps.get_model("passbook_flows", "FlowStageBinding")
    PolicyBinding = apps.get_model("passbook_policies", "PolicyBinding")

    ExpressionPolicy = apps.get_model(
        "passbook_policies_expression", "ExpressionPolicy"
    )

    PromptStage = apps.get_model("passbook_stages_prompt", "PromptStage")
    Prompt = apps.get_model("passbook_stages_prompt", "Prompt")
    UserWriteStage = apps.get_model("passbook_stages_user_write", "UserWriteStage")
    UserLoginStage = apps.get_model("passbook_stages_user_login", "UserLoginStage")

    db_alias = schema_editor.connection.alias

    # Create a policy that only allows this flow when doing an SSO Request
    flow_policy = ExpressionPolicy.objects.using(db_alias).create(
        name="default-source-enrollment-if-sso", expression=FLOW_POLICY_EXPRESSION
    )

    # This creates a Flow used by sources to enroll users
    # It makes sure that a username is set, and if not, prompts the user for a Username
    flow = Flow.objects.using(db_alias).create(
        name="default-source-enrollment",
        slug="default-source-enrollment",
        designation=FlowDesignation.ENROLLMENT,
    )
    PolicyBinding.objects.using(db_alias).create(
        policy=flow_policy, target=flow, order=0
    )

    # PromptStage to ask user for their username
    prompt_stage = PromptStage.objects.using(db_alias).create(
        name="default-source-enrollment-username-prompt",
    )
    prompt_stage.fields.add(
        Prompt.objects.using(db_alias).create(
            field_key="username",
            label="Username",
            type=FieldTypes.TEXT,
            required=True,
            placeholder="Username",
        )
    )
    # Policy to only trigger prompt when no username is given
    prompt_policy = ExpressionPolicy.objects.using(db_alias).create(
        name="default-source-enrollment-if-username",
        expression=PROMPT_POLICY_EXPRESSION,
    )

    # UserWrite stage to create the user, and login stage to log user in
    user_write = UserWriteStage.objects.using(db_alias).create(
        name="default-source-enrollment-write"
    )
    user_login = UserLoginStage.objects.using(db_alias).create(
        name="default-source-enrollment-login"
    )

    binding = FlowStageBinding.objects.using(db_alias).create(
        flow=flow, stage=prompt_stage, order=0
    )
    PolicyBinding.objects.using(db_alias).create(
        policy=prompt_policy, target=binding, order=0
    )

    FlowStageBinding.objects.using(db_alias).create(
        flow=flow, stage=user_write, order=1
    )
    FlowStageBinding.objects.using(db_alias).create(
        flow=flow, stage=user_login, order=2
    )


def create_default_source_authentication_flow(
    apps: Apps, schema_editor: BaseDatabaseSchemaEditor
):
    Flow = apps.get_model("passbook_flows", "Flow")
    FlowStageBinding = apps.get_model("passbook_flows", "FlowStageBinding")
    PolicyBinding = apps.get_model("passbook_policies", "PolicyBinding")

    ExpressionPolicy = apps.get_model(
        "passbook_policies_expression", "ExpressionPolicy"
    )

    UserLoginStage = apps.get_model("passbook_stages_user_login", "UserLoginStage")

    db_alias = schema_editor.connection.alias

    # Create a policy that only allows this flow when doing an SSO Request
    flow_policy = ExpressionPolicy.objects.using(db_alias).create(
        name="default-source-authentication-if-sso", expression=FLOW_POLICY_EXPRESSION
    )

    # This creates a Flow used by sources to authenticate users
    flow = Flow.objects.using(db_alias).create(
        name="default-source-authentication",
        slug="default-source-authentication",
        designation=FlowDesignation.AUTHENTICATION,
    )
    PolicyBinding.objects.using(db_alias).create(
        policy=flow_policy, target=flow, order=0
    )

    user_login = UserLoginStage.objects.using(db_alias).create(
        name="default-source-authentication-login"
    )
    FlowStageBinding.objects.using(db_alias).create(
        flow=flow, stage=user_login, order=0
    )


class Migration(migrations.Migration):

    dependencies = [
        ("passbook_flows", "0003_auto_20200523_1133"),
        ("passbook_policies", "0001_initial"),
        ("passbook_policies_expression", "0001_initial"),
        ("passbook_stages_prompt", "0001_initial"),
        ("passbook_stages_user_write", "0001_initial"),
        ("passbook_stages_user_login", "0001_initial"),
    ]

    operations = [
        migrations.RunPython(create_default_source_enrollment_flow),
        migrations.RunPython(create_default_source_authentication_flow),
    ]
