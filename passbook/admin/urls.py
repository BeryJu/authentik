"""passbook URL Configuration"""
from django.urls import path

from passbook.admin.views import (
    applications,
    certificate_key_pair,
    flows,
    groups,
    outposts,
    overview,
    policies,
    policies_bindings,
    property_mapping,
    providers,
    sources,
    stages,
    stages_bindings,
    stages_invitations,
    stages_prompts,
    tokens,
    users,
)

urlpatterns = [
    path("", overview.AdministrationOverviewView.as_view(), name="overview"),
    # Applications
    path(
        "applications/", applications.ApplicationListView.as_view(), name="applications"
    ),
    path(
        "applications/create/",
        applications.ApplicationCreateView.as_view(),
        name="application-create",
    ),
    path(
        "applications/<uuid:pk>/update/",
        applications.ApplicationUpdateView.as_view(),
        name="application-update",
    ),
    path(
        "applications/<uuid:pk>/delete/",
        applications.ApplicationDeleteView.as_view(),
        name="application-delete",
    ),
    # Tokens
    path("tokens/", tokens.TokenListView.as_view(), name="tokens"),
    path(
        "tokens/<uuid:pk>/delete/",
        tokens.TokenDeleteView.as_view(),
        name="token-delete",
    ),
    # Sources
    path("sources/", sources.SourceListView.as_view(), name="sources"),
    path("sources/create/", sources.SourceCreateView.as_view(), name="source-create"),
    path(
        "sources/<uuid:pk>/update/",
        sources.SourceUpdateView.as_view(),
        name="source-update",
    ),
    path(
        "sources/<uuid:pk>/delete/",
        sources.SourceDeleteView.as_view(),
        name="source-delete",
    ),
    # Policies
    path("policies/", policies.PolicyListView.as_view(), name="policies"),
    path("policies/create/", policies.PolicyCreateView.as_view(), name="policy-create"),
    path(
        "policies/<uuid:pk>/update/",
        policies.PolicyUpdateView.as_view(),
        name="policy-update",
    ),
    path(
        "policies/<uuid:pk>/delete/",
        policies.PolicyDeleteView.as_view(),
        name="policy-delete",
    ),
    path(
        "policies/<uuid:pk>/test/",
        policies.PolicyTestView.as_view(),
        name="policy-test",
    ),
    # Policy bindings
    path(
        "policies/bindings/",
        policies_bindings.PolicyBindingListView.as_view(),
        name="policies-bindings",
    ),
    path(
        "policies/bindings/create/",
        policies_bindings.PolicyBindingCreateView.as_view(),
        name="policy-binding-create",
    ),
    path(
        "policies/bindings/<uuid:pk>/update/",
        policies_bindings.PolicyBindingUpdateView.as_view(),
        name="policy-binding-update",
    ),
    path(
        "policies/bindings/<uuid:pk>/delete/",
        policies_bindings.PolicyBindingDeleteView.as_view(),
        name="policy-binding-delete",
    ),
    # Providers
    path("providers/", providers.ProviderListView.as_view(), name="providers"),
    path(
        "providers/create/",
        providers.ProviderCreateView.as_view(),
        name="provider-create",
    ),
    path(
        "providers/<int:pk>/update/",
        providers.ProviderUpdateView.as_view(),
        name="provider-update",
    ),
    path(
        "providers/<int:pk>/delete/",
        providers.ProviderDeleteView.as_view(),
        name="provider-delete",
    ),
    # Stages
    path("stages/", stages.StageListView.as_view(), name="stages"),
    path("stages/create/", stages.StageCreateView.as_view(), name="stage-create"),
    path(
        "stages/<uuid:pk>/update/",
        stages.StageUpdateView.as_view(),
        name="stage-update",
    ),
    path(
        "stages/<uuid:pk>/delete/",
        stages.StageDeleteView.as_view(),
        name="stage-delete",
    ),
    # Stage bindings
    path(
        "stages/bindings/",
        stages_bindings.StageBindingListView.as_view(),
        name="stage-bindings",
    ),
    path(
        "stages/bindings/create/",
        stages_bindings.StageBindingCreateView.as_view(),
        name="stage-binding-create",
    ),
    path(
        "stages/bindings/<uuid:pk>/update/",
        stages_bindings.StageBindingUpdateView.as_view(),
        name="stage-binding-update",
    ),
    path(
        "stages/bindings/<uuid:pk>/delete/",
        stages_bindings.StageBindingDeleteView.as_view(),
        name="stage-binding-delete",
    ),
    # Stage Prompts
    path(
        "stages/prompts/",
        stages_prompts.PromptListView.as_view(),
        name="stage-prompts",
    ),
    path(
        "stages/prompts/create/",
        stages_prompts.PromptCreateView.as_view(),
        name="stage-prompt-create",
    ),
    path(
        "stages/prompts/<uuid:pk>/update/",
        stages_prompts.PromptUpdateView.as_view(),
        name="stage-prompt-update",
    ),
    path(
        "stages/prompts/<uuid:pk>/delete/",
        stages_prompts.PromptDeleteView.as_view(),
        name="stage-prompt-delete",
    ),
    # Stage Invitations
    path(
        "stages/invitations/",
        stages_invitations.InvitationListView.as_view(),
        name="stage-invitations",
    ),
    path(
        "stages/invitations/create/",
        stages_invitations.InvitationCreateView.as_view(),
        name="stage-invitation-create",
    ),
    path(
        "stages/invitations/<uuid:pk>/delete/",
        stages_invitations.InvitationDeleteView.as_view(),
        name="stage-invitation-delete",
    ),
    # Flows
    path("flows/", flows.FlowListView.as_view(), name="flows"),
    path("flows/create/", flows.FlowCreateView.as_view(), name="flow-create",),
    path("flows/import/", flows.FlowImportView.as_view(), name="flow-import",),
    path(
        "flows/<uuid:pk>/update/", flows.FlowUpdateView.as_view(), name="flow-update",
    ),
    path(
        "flows/<uuid:pk>/execute/",
        flows.FlowDebugExecuteView.as_view(),
        name="flow-execute",
    ),
    path(
        "flows/<uuid:pk>/export/", flows.FlowExportView.as_view(), name="flow-export",
    ),
    path(
        "flows/<uuid:pk>/delete/", flows.FlowDeleteView.as_view(), name="flow-delete",
    ),
    # Property Mappings
    path(
        "property-mappings/",
        property_mapping.PropertyMappingListView.as_view(),
        name="property-mappings",
    ),
    path(
        "property-mappings/create/",
        property_mapping.PropertyMappingCreateView.as_view(),
        name="property-mapping-create",
    ),
    path(
        "property-mappings/<uuid:pk>/update/",
        property_mapping.PropertyMappingUpdateView.as_view(),
        name="property-mapping-update",
    ),
    path(
        "property-mappings/<uuid:pk>/delete/",
        property_mapping.PropertyMappingDeleteView.as_view(),
        name="property-mapping-delete",
    ),
    # Users
    path("users/", users.UserListView.as_view(), name="users"),
    path("users/create/", users.UserCreateView.as_view(), name="user-create"),
    path("users/<int:pk>/update/", users.UserUpdateView.as_view(), name="user-update"),
    path("users/<int:pk>/delete/", users.UserDeleteView.as_view(), name="user-delete"),
    path(
        "users/<int:pk>/reset/",
        users.UserPasswordResetView.as_view(),
        name="user-password-reset",
    ),
    # Groups
    path("groups/", groups.GroupListView.as_view(), name="groups"),
    path("groups/create/", groups.GroupCreateView.as_view(), name="group-create"),
    path(
        "groups/<uuid:pk>/update/",
        groups.GroupUpdateView.as_view(),
        name="group-update",
    ),
    path(
        "groups/<uuid:pk>/delete/",
        groups.GroupDeleteView.as_view(),
        name="group-delete",
    ),
    # Certificate-Key Pairs
    path(
        "crypto/certificates/",
        certificate_key_pair.CertificateKeyPairListView.as_view(),
        name="certificate_key_pair",
    ),
    path(
        "crypto/certificates/create/",
        certificate_key_pair.CertificateKeyPairCreateView.as_view(),
        name="certificatekeypair-create",
    ),
    path(
        "crypto/certificates/<uuid:pk>/update/",
        certificate_key_pair.CertificateKeyPairUpdateView.as_view(),
        name="certificatekeypair-update",
    ),
    path(
        "crypto/certificates/<uuid:pk>/delete/",
        certificate_key_pair.CertificateKeyPairDeleteView.as_view(),
        name="certificatekeypair-delete",
    ),
    # Outposts
    path("outposts/", outposts.OutpostListView.as_view(), name="outposts",),
    path(
        "outposts/create/", outposts.OutpostCreateView.as_view(), name="outpost-create",
    ),
    path(
        "outposts/<uuid:pk>/update/",
        outposts.OutpostUpdateView.as_view(),
        name="outpost-update",
    ),
    path(
        "outposts/<uuid:pk>/delete/",
        outposts.OutpostDeleteView.as_view(),
        name="outpost-delete",
    ),
]
