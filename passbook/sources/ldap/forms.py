"""passbook LDAP Forms"""

from django import forms
from django.contrib.admin.widgets import FilteredSelectMultiple
from django.utils.translation import gettext_lazy as _

from passbook.admin.fields import CodeMirrorWidget
from passbook.core.expression import PropertyMappingEvaluator
from passbook.sources.ldap.models import LDAPPropertyMapping, LDAPSource


class LDAPSourceForm(forms.ModelForm):
    """LDAPSource Form"""

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.fields["property_mappings"].queryset = LDAPPropertyMapping.objects.all()

    class Meta:

        model = LDAPSource
        fields = [
            # we don't use all common fields, as we don't use flows for this
            "name",
            "slug",
            "enabled",
            # -- start of our custom fields
            "server_uri",
            "bind_cn",
            "bind_password",
            "start_tls",
            "base_dn",
            "additional_user_dn",
            "additional_group_dn",
            "user_object_filter",
            "group_object_filter",
            "user_group_membership_field",
            "object_uniqueness_field",
            "sync_users",
            "sync_groups",
            "sync_parent_group",
            "property_mappings",
        ]
        widgets = {
            "name": forms.TextInput(),
            "server_uri": forms.TextInput(),
            "bind_cn": forms.TextInput(),
            "bind_password": forms.TextInput(),
            "base_dn": forms.TextInput(),
            "additional_user_dn": forms.TextInput(),
            "additional_group_dn": forms.TextInput(),
            "user_object_filter": forms.TextInput(),
            "group_object_filter": forms.TextInput(),
            "user_group_membership_field": forms.TextInput(),
            "object_uniqueness_field": forms.TextInput(),
            "property_mappings": FilteredSelectMultiple(_("Property Mappings"), False),
        }


class LDAPPropertyMappingForm(forms.ModelForm):
    """LDAP Property Mapping form"""

    template_name = "ldap/property_mapping_form.html"

    def clean_expression(self):
        """Test Syntax"""
        expression = self.cleaned_data.get("expression")
        evaluator = PropertyMappingEvaluator()
        evaluator.validate(expression)
        return expression

    class Meta:

        model = LDAPPropertyMapping
        fields = ["name", "object_field", "expression"]
        widgets = {
            "name": forms.TextInput(),
            "ldap_property": forms.TextInput(),
            "object_field": forms.TextInput(),
            "expression": CodeMirrorWidget(mode="python"),
        }
        help_texts = {
            "object_field": _("Field of the user object this value is written to.")
        }
