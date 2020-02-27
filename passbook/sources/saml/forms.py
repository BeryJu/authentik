"""passbook SAML SP Forms"""

from django import forms
from django.contrib.admin.widgets import FilteredSelectMultiple
from django.utils.translation import gettext as _

from passbook.admin.forms.source import SOURCE_FORM_FIELDS
from passbook.providers.saml.utils.cert import CertificateBuilder
from passbook.sources.saml.models import SAMLSource


class SAMLSourceForm(forms.ModelForm):
    """SAML Provider form"""

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        builder = CertificateBuilder()
        builder.build()
        self.fields["signing_cert"].initial = builder.certificate

    class Meta:

        model = SAMLSource
        fields = SOURCE_FORM_FIELDS + [
            "issuer",
            "idp_url",
            "idp_logout_url",
            "auto_logout",
            "signing_cert",
        ]
        widgets = {
            "name": forms.TextInput(),
            "policies": FilteredSelectMultiple(_("policies"), False),
            "issuer": forms.TextInput(),
            "idp_url": forms.TextInput(),
            "idp_logout_url": forms.TextInput(),
        }
