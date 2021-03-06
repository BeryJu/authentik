import { ConsentStage, ConsentStageModeEnum, StagesApi } from "authentik-api";
import { t } from "@lingui/macro";
import { customElement, property } from "lit-element";
import { html, TemplateResult } from "lit-html";
import { DEFAULT_CONFIG } from "../../../api/Config";
import { Form } from "../../../elements/forms/Form";
import { ifDefined } from "lit-html/directives/if-defined";
import "../../../elements/forms/HorizontalFormElement";
import "../../../elements/forms/FormGroup";

@customElement("ak-stage-consent-form")
export class ConsentStageForm extends Form<ConsentStage> {

    set stageUUID(value: string) {
        new StagesApi(DEFAULT_CONFIG).stagesConsentRead({
            stageUuid: value,
        }).then(stage => {
            this.stage = stage;
            this.showExpiresIn = stage.name === ConsentStageModeEnum.Expiring;
        });
    }

    @property({attribute: false})
    stage?: ConsentStage;

    @property({type: Boolean})
    showExpiresIn = false;

    getSuccessMessage(): string {
        if (this.stage) {
            return t`Successfully updated stage.`;
        } else {
            return t`Successfully created stage.`;
        }
    }

    send = (data: ConsentStage): Promise<ConsentStage> => {
        if (this.stage) {
            return new StagesApi(DEFAULT_CONFIG).stagesConsentUpdate({
                stageUuid: this.stage.pk || "",
                data: data
            });
        } else {
            return new StagesApi(DEFAULT_CONFIG).stagesConsentCreate({
                data: data
            });
        }
    };

    renderForm(): TemplateResult {
        return html`<form class="pf-c-form pf-m-horizontal">
            <div class="form-help-text">
                ${t`Prompt for the user's consent. The consent can either be permanent or expire in a defined amount of time.`}
            </div>
            <ak-form-element-horizontal
                label=${t`Name`}
                ?required=${true}
                name="name">
                <input type="text" value="${ifDefined(this.stage?.name || "")}" class="pf-c-form-control" required>
            </ak-form-element-horizontal>
            <ak-form-group .expanded=${true}>
                <span slot="header">
                    ${t`Stage-specific settings`}
                </span>
                <div slot="body" class="pf-c-form">
                    <ak-form-element-horizontal
                        label=${t`Mode`}
                        ?required=${true}
                        name="mode">
                        <select class="pf-c-form-control" @change=${(ev: Event) => {
                            const target = ev.target as HTMLSelectElement;
                            if (target.selectedOptions[0].value === ConsentStageModeEnum.Expiring) {
                                this.showExpiresIn = true;
                            } else {
                                this.showExpiresIn = false;
                            }
                        }}>
                            <option value=${ConsentStageModeEnum.AlwaysRequire} ?selected=${this.stage?.mode === ConsentStageModeEnum.AlwaysRequire}>
                                ${t`Always require consent`}
                            </option>
                            <option value=${ConsentStageModeEnum.Permanent} ?selected=${this.stage?.mode === ConsentStageModeEnum.Permanent}>
                                ${t`Consent given last indefinitely`}
                            </option>
                            <option value=${ConsentStageModeEnum.Expiring} ?selected=${this.stage?.mode === ConsentStageModeEnum.Expiring}>
                                ${t`Consent expires.`}
                            </option>
                        </select>
                    </ak-form-element-horizontal>
                    <ak-form-element-horizontal
                        ?hidden=${!this.showExpiresIn}
                        label=${t`Consent expires in`}
                        ?required=${true}
                        name="consentExpireIn">
                        <input type="text" value="${ifDefined(this.stage?.consentExpireIn || "weeks=4")}" class="pf-c-form-control" required>
                        <p class="pf-c-form__helper-text">${t`Offset after which consent expires. (Format: hours=1;minutes=2;seconds=3).`}</p>
                    </ak-form-element-horizontal>
                </div>
            </ak-form-group>
        </form>`;
    }

}
