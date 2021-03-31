"""authentik SAML SP Forms"""

from django import forms
from django.utils.translation import gettext_lazy as _

from authentik.crypto.models import CertificateKeyPair
from authentik.flows.models import Flow, FlowDesignation
from authentik.sources.saml.models import SAMLSource


class SAMLSourceForm(forms.ModelForm):
    """SAML Provider form"""

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

        self.fields["pre_authentication_flow"].queryset = Flow.objects.filter(
            designation=FlowDesignation.STAGE_CONFIGURATION
        )
        self.fields["authentication_flow"].queryset = Flow.objects.filter(
            designation=FlowDesignation.AUTHENTICATION
        )
        self.fields["enrollment_flow"].queryset = Flow.objects.filter(
            designation=FlowDesignation.ENROLLMENT
        )
        self.fields["signing_kp"].queryset = CertificateKeyPair.objects.filter(
            certificate_data__isnull=False,
            key_data__isnull=False,
        )

    class Meta:

        model = SAMLSource
        fields = [
            "name",
            "slug",
            "enabled",
            "policy_engine_mode",
            "pre_authentication_flow",
            "authentication_flow",
            "enrollment_flow",
            "issuer",
            "sso_url",
            "slo_url",
            "binding_type",
            "name_id_policy",
            "allow_idp_initiated",
            "signing_kp",
            "digest_algorithm",
            "signature_algorithm",
            "temporary_user_delete_after",
        ]
        widgets = {
            "name": forms.TextInput(),
            "issuer": forms.TextInput(),
            "sso_url": forms.TextInput(),
            "slo_url": forms.TextInput(),
            "temporary_user_delete_after": forms.TextInput(),
        }
        labels = {
            "name_id_policy": _("Name ID Policy"),
            "allow_idp_initiated": _("Allow IDP-initiated logins"),
        }
