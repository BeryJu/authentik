"""passbook core authentication forms"""
from django import forms
from django.core.exceptions import ValidationError
from django.core.validators import validate_email
from django.utils.translation import gettext_lazy as _
from structlog import get_logger

from passbook.core.models import User
from passbook.lib.config import CONFIG
from passbook.lib.utils.ui import human_list

LOGGER = get_logger()


class LoginForm(forms.Form):
    """Allow users to login"""

    title = _("Log in to your account")
    uid_field = forms.CharField(label=_(""))

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        if CONFIG.y("passbook.uid_fields") == ["e-mail"]:
            self.fields["uid_field"] = forms.EmailField()
        self.fields["uid_field"].label = human_list(
            [x.title() for x in CONFIG.y("passbook.uid_fields")]
        )

    def clean_uid_field(self):
        """Validate uid_field after EmailValidator if 'email' is the only selected uid_fields"""
        if CONFIG.y("passbook.uid_fields") == ["email"]:
            validate_email(self.cleaned_data.get("uid_field"))
        return self.cleaned_data.get("uid_field")


class SignUpForm(forms.Form):
    """SignUp Form"""

    title = _("Sign Up")
    name = forms.CharField(
        label=_("Name"), widget=forms.TextInput(attrs={"placeholder": _("Name")})
    )
    username = forms.CharField(
        label=_("Username"),
        widget=forms.TextInput(attrs={"placeholder": _("Username")}),
    )
    email = forms.EmailField(
        label=_("E-Mail"), widget=forms.TextInput(attrs={"placeholder": _("E-Mail")})
    )
    password = forms.CharField(
        label=_("Password"),
        widget=forms.PasswordInput(attrs={"placeholder": _("Password")}),
    )
    password_repeat = forms.CharField(
        label=_("Repeat Password"),
        widget=forms.PasswordInput(attrs={"placeholder": _("Repeat Password")}),
    )

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        # All fields which have initial data supplied are set to read only
        if "initial" in kwargs:
            for field in kwargs.get("initial").keys():
                self.fields[field].widget.attrs["readonly"] = "readonly"

    def clean_username(self):
        """Check if username is used already"""
        username = self.cleaned_data.get("username")
        if User.objects.filter(username=username).exists():
            LOGGER.warning("username already exists", username=username)
            raise ValidationError(_("Username already exists"))
        return username

    def clean_email(self):
        """Check if email is already used in django or other auth sources"""
        email = self.cleaned_data.get("email")
        # Check if user exists already, error early
        if User.objects.filter(email=email).exists():
            LOGGER.debug("email already exists", email=email)
            raise ValidationError(_("Email already exists"))
        return email

    def clean_password_repeat(self):
        """Check if Password adheres to filter and if passwords matche"""
        password = self.cleaned_data.get("password")
        password_repeat = self.cleaned_data.get("password_repeat")
        if password != password_repeat:
            raise ValidationError(_("Passwords don't match"))
        return self.cleaned_data.get("password_repeat")
