import { AuthenticateWebAuthnStage, StagesApi } from "authentik-api";
import { t } from "@lingui/macro";
import { customElement, property } from "lit-element";
import { html, TemplateResult } from "lit-html";
import { DEFAULT_CONFIG } from "../../../api/Config";
import { Form } from "../../../elements/forms/Form";
import { ifDefined } from "lit-html/directives/if-defined";
import "../../../elements/forms/HorizontalFormElement";

@customElement("ak-stage-authenticator-webauthn-form")
export class AuthenticateWebAuthnStageForm extends Form<AuthenticateWebAuthnStage> {

    set stageUUID(value: string) {
        new StagesApi(DEFAULT_CONFIG).stagesAuthenticatorWebauthnRead({
            stageUuid: value,
        }).then(stage => {
            this.stage = stage;
        });
    }

    @property({attribute: false})
    stage?: AuthenticateWebAuthnStage;

    getSuccessMessage(): string {
        if (this.stage) {
            return t`Successfully updated stage.`;
        } else {
            return t`Successfully created stage.`;
        }
    }

    send = (data: AuthenticateWebAuthnStage): Promise<AuthenticateWebAuthnStage> => {
        if (this.stage) {
            return new StagesApi(DEFAULT_CONFIG).stagesAuthenticatorWebauthnUpdate({
                stageUuid: this.stage.pk || "",
                data: data
            });
        } else {
            return new StagesApi(DEFAULT_CONFIG).stagesAuthenticatorWebauthnCreate({
                data: data
            });
        }
    };

    renderForm(): TemplateResult {
        return html`<form class="pf-c-form pf-m-horizontal">
            <div class="form-help-text">
                ${t`Stage used to configure a WebAutnn authenticator (i.e. Yubikey, FaceID/Windows Hello).`}
            </div>
            <ak-form-element-horizontal
                label=${t`Name`}
                ?required=${true}
                name="name">
                <input type="text" value="${ifDefined(this.stage?.name || "")}" class="pf-c-form-control" required>
            </ak-form-element-horizontal>
        </form>`;
    }

}
