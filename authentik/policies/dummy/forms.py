"""authentik Policy forms"""

from django import forms
from django.utils.translation import gettext as _

from authentik.policies.dummy.models import DummyPolicy
from authentik.policies.forms import GENERAL_FIELDS


class DummyPolicyForm(forms.ModelForm):
    """DummyPolicyForm Form"""

    class Meta:

        model = DummyPolicy
        fields = GENERAL_FIELDS + ["result", "wait_min", "wait_max"]
        widgets = {
            "name": forms.TextInput(),
        }
        labels = {"result": _("Allow user")}
