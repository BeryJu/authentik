import { UserLoginStage, StagesApi } from "authentik-api";
import { t } from "@lingui/macro";
import { customElement, property } from "lit-element";
import { html, TemplateResult } from "lit-html";
import { DEFAULT_CONFIG } from "../../../api/Config";
import { Form } from "../../../elements/forms/Form";
import "../../../elements/forms/HorizontalFormElement";
import "../../../elements/forms/FormGroup";
import { first } from "../../../utils";

@customElement("ak-stage-user-login-form")
export class UserLoginStageForm extends Form<UserLoginStage> {

    set stageUUID(value: string) {
        new StagesApi(DEFAULT_CONFIG).stagesUserLoginRead({
            stageUuid: value,
        }).then(stage => {
            this.stage = stage;
        });
    }

    @property({attribute: false})
    stage?: UserLoginStage;

    getSuccessMessage(): string {
        if (this.stage) {
            return t`Successfully updated stage.`;
        } else {
            return t`Successfully created stage.`;
        }
    }

    send = (data: UserLoginStage): Promise<UserLoginStage> => {
        if (this.stage) {
            return new StagesApi(DEFAULT_CONFIG).stagesUserLoginUpdate({
                stageUuid: this.stage.pk || "",
                data: data
            });
        } else {
            return new StagesApi(DEFAULT_CONFIG).stagesUserLoginCreate({
                data: data
            });
        }
    };

    renderForm(): TemplateResult {
        return html`<form class="pf-c-form pf-m-horizontal">
            <div class="form-help-text">
                ${t`Log the currently pending user in.`}
            </div>
            <ak-form-element-horizontal
                label=${t`Name`}
                ?required=${true}
                name="name">
                <input type="text" value="${first(this.stage?.name, "")}" class="pf-c-form-control" required>
            </ak-form-element-horizontal>
            <ak-form-group .expanded=${true}>
                <span slot="header">
                    ${t`Stage-specific settings`}
                </span>
                <div slot="body" class="pf-c-form">
                    <ak-form-element-horizontal
                        label=${t`Session duration`}
                        ?required=${true}
                        name="sessionDuration">
                        <input type="text" value="${first(this.stage?.sessionDuration, "seconds=0")}" class="pf-c-form-control" required>
                        <p class="pf-c-form__helper-text">${t`Determines how long a session lasts. Default of 0 seconds means that the sessions lasts until the browser is closed.`}</p>
                        <p class="pf-c-form__helper-text">${t`(Format: hours=-1;minutes=-2;seconds=-3).`}</p>
                    </ak-form-element-horizontal>
                </div>
            </ak-form-group>
        </form>`;
    }

}
