"""Validation stage challenge checking"""
from django.http import HttpRequest
from django.utils.translation import gettext_lazy as _
from django_otp import match_token
from django_otp.models import Device
from django_otp.plugins.otp_static.models import StaticDevice
from django_otp.plugins.otp_totp.models import TOTPDevice
from rest_framework.fields import CharField, JSONField
from rest_framework.serializers import ValidationError
from webauthn import WebAuthnAssertionOptions, WebAuthnAssertionResponse, WebAuthnUser
from webauthn.webauthn import (
    AuthenticationRejectedException,
    RegistrationRejectedException,
    WebAuthnUserDataMissing,
)

from authentik.core.api.utils import PassiveSerializer
from authentik.core.models import User
from authentik.stages.authenticator_webauthn.models import WebAuthnDevice
from authentik.stages.authenticator_webauthn.utils import generate_challenge, get_origin


class DeviceChallenge(PassiveSerializer):
    """Single device challenge"""

    device_class = CharField()
    device_uid = CharField()
    challenge = JSONField()


def get_challenge_for_device(request: HttpRequest, device: Device) -> dict:
    """Generate challenge for a single device"""
    if isinstance(device, (TOTPDevice, StaticDevice)):
        # Code-based challenges have no hints
        return {}
    return get_webauthn_challenge(request, device)


def get_webauthn_challenge(request: HttpRequest, device: WebAuthnDevice) -> dict:
    """Send the client a challenge that we'll check later"""
    request.session.pop("challenge", None)

    challenge = generate_challenge(32)

    # We strip the padding from the challenge stored in the session
    # for the reasons outlined in the comment in webauthn_begin_activate.
    request.session["challenge"] = challenge.rstrip("=")

    webauthn_user = WebAuthnUser(
        device.user.uid,
        device.user.username,
        device.user.name,
        device.user.avatar,
        device.credential_id,
        device.public_key,
        device.sign_count,
        device.rp_id,
    )

    webauthn_assertion_options = WebAuthnAssertionOptions(webauthn_user, challenge)

    return webauthn_assertion_options.assertion_dict


def validate_challenge_code(code: str, request: HttpRequest, user: User) -> str:
    """Validate code-based challenges. We test against every device, on purpose, as
    the user mustn't choose between totp and static devices."""
    device = match_token(user, code)
    if not device:
        raise ValidationError(_("Invalid Token"))
    return code


def validate_challenge_webauthn(data: dict, request: HttpRequest, user: User) -> dict:
    """Validate WebAuthn Challenge"""
    challenge = request.session.get("challenge")
    assertion_response = data
    credential_id = assertion_response.get("id")

    device = WebAuthnDevice.objects.filter(credential_id=credential_id).first()
    if not device:
        raise ValidationError("Device does not exist.")

    webauthn_user = WebAuthnUser(
        user.uid,
        user.username,
        user.name,
        user.avatar,
        device.credential_id,
        device.public_key,
        device.sign_count,
        device.rp_id,
    )

    webauthn_assertion_response = WebAuthnAssertionResponse(
        webauthn_user,
        assertion_response,
        challenge,
        get_origin(request),
        uv_required=False,
    )  # User Verification

    try:
        sign_count = webauthn_assertion_response.verify()
    except (
        AuthenticationRejectedException,
        WebAuthnUserDataMissing,
        RegistrationRejectedException,
    ) as exc:
        raise ValidationError("Assertion failed") from exc

    device.set_sign_count(sign_count)
    return data
