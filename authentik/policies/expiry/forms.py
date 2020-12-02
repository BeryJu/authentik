"""authentik PasswordExpiry Policy forms"""

from django import forms
from django.utils.translation import gettext as _

from authentik.policies.expiry.models import PasswordExpiryPolicy
from authentik.policies.forms import GENERAL_FIELDS


class PasswordExpiryPolicyForm(forms.ModelForm):
    """Edit PasswordExpiryPolicy instances"""

    class Meta:

        model = PasswordExpiryPolicy
        fields = GENERAL_FIELDS + ["days", "deny_only"]
        widgets = {
            "name": forms.TextInput(),
            "order": forms.NumberInput(),
            "days": forms.NumberInput(),
        }
        labels = {"deny_only": _("Only fail the policy, don't set user's password.")}
