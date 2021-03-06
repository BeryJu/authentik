import { AuthenticatorValidateStage, AuthenticatorValidateStageNotConfiguredActionEnum, AuthenticatorValidateStageDeviceClassesEnum, StagesApi } from "authentik-api";
import { t } from "@lingui/macro";
import { customElement, property } from "lit-element";
import { html, TemplateResult } from "lit-html";
import { DEFAULT_CONFIG } from "../../../api/Config";
import { Form } from "../../../elements/forms/Form";
import { ifDefined } from "lit-html/directives/if-defined";
import "../../../elements/forms/HorizontalFormElement";
import "../../../elements/forms/FormGroup";
import { until } from "lit-html/directives/until";

@customElement("ak-stage-authenticator-validate-form")
export class AuthenticatorValidateStageForm extends Form<AuthenticatorValidateStage> {

    set stageUUID(value: string) {
        new StagesApi(DEFAULT_CONFIG).stagesAuthenticatorValidateRead({
            stageUuid: value,
        }).then(stage => {
            this.stage = stage;
            this.showConfigureFlow = stage.notConfiguredAction === AuthenticatorValidateStageNotConfiguredActionEnum.Configure;
        });
    }

    @property({attribute: false})
    stage?: AuthenticatorValidateStage;

    @property({ type: Boolean })
    showConfigureFlow = false;

    getSuccessMessage(): string {
        if (this.stage) {
            return t`Successfully updated stage.`;
        } else {
            return t`Successfully created stage.`;
        }
    }

    send = (data: AuthenticatorValidateStage): Promise<AuthenticatorValidateStage> => {
        if (this.stage) {
            return new StagesApi(DEFAULT_CONFIG).stagesAuthenticatorValidateUpdate({
                stageUuid: this.stage.pk || "",
                data: data
            });
        } else {
            return new StagesApi(DEFAULT_CONFIG).stagesAuthenticatorValidateCreate({
                data: data
            });
        }
    };

    isDeviceClassSelected(field: AuthenticatorValidateStageDeviceClassesEnum): boolean {
        return (this.stage?.deviceClasses || []).filter(isField => {
            return field === isField;
        }).length > 0;
    }

    renderForm(): TemplateResult {
        return html`<form class="pf-c-form pf-m-horizontal">
            <div class="form-help-text">
                ${t`Stage used to validate any authenticator. This stage should be used during authentication or authorization flows.`}
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
                        label=${t`Not configured action`}
                        ?required=${true}
                        name="mode">
                        <select class="pf-c-form-control" @change=${(ev: Event) => {
                            const target = ev.target as HTMLSelectElement;
                            if (target.selectedOptions[0].value === AuthenticatorValidateStageNotConfiguredActionEnum.Configure) {
                                this.showConfigureFlow = true;
                            } else {
                                this.showConfigureFlow = false;
                            }
                        }}>
                            <option value=${AuthenticatorValidateStageNotConfiguredActionEnum.Configure} ?selected=${this.stage?.notConfiguredAction === AuthenticatorValidateStageNotConfiguredActionEnum.Configure}>
                                ${t`Force the user to configure an authenticator`}
                            </option>
                            <option value=${AuthenticatorValidateStageNotConfiguredActionEnum.Deny} ?selected=${this.stage?.notConfiguredAction === AuthenticatorValidateStageNotConfiguredActionEnum.Deny}>
                                ${t`Deny the user access`}
                            </option>
                            <option value=${AuthenticatorValidateStageNotConfiguredActionEnum.Skip} ?selected=${this.stage?.notConfiguredAction === AuthenticatorValidateStageNotConfiguredActionEnum.Skip}>
                                ${t`Continue`}
                            </option>
                        </select>
                    </ak-form-element-horizontal>
                    <ak-form-element-horizontal
                        label=${t`User fields`}
                        ?required=${true}
                        name="transports">
                        <select name="users" class="pf-c-form-control" multiple>
                            <option value=${AuthenticatorValidateStageDeviceClassesEnum.Static} ?selected=${this.isDeviceClassSelected(AuthenticatorValidateStageDeviceClassesEnum.Static)}>
                                ${t`Static Tokens`}
                            </option>
                            <option value=${AuthenticatorValidateStageDeviceClassesEnum.Totp} ?selected=${this.isDeviceClassSelected(AuthenticatorValidateStageDeviceClassesEnum.Totp)}>
                                ${t`TOTP Authenticators`}
                            </option>
                            <option value=${AuthenticatorValidateStageDeviceClassesEnum.Webauthn} ?selected=${this.isDeviceClassSelected(AuthenticatorValidateStageDeviceClassesEnum.Webauthn)}>
                                ${t`WebAuthn Authenticators`}
                            </option>
                        </select>
                        <p class="pf-c-form__helper-text">${t`Device classes which can be used to authenticate.`}</p>
                        <p class="pf-c-form__helper-text">${t`Hold control/command to select multiple items.`}</p>
                    </ak-form-element-horizontal>
                    ${this.showConfigureFlow ? html`
                    <ak-form-element-horizontal
                        label=${t`Configuration flow`}
                        ?required=${true}
                        name="configureFlow">
                        <select class="pf-c-form-control">
                            <option value="" ?selected=${this.stage?.configurationStage === undefined}>---------</option>
                            ${until(new StagesApi(DEFAULT_CONFIG).stagesAllList({
                                ordering: "pk",
                            }).then(stages => {
                                return stages.results.map(stage => {
                                    const selected = this.stage?.configurationStage === stage.pk;
                                    return html`<option value=${ifDefined(stage.pk)} ?selected=${selected}>${stage.name} (${stage.verboseName})</option>`;
                                });
                            }), html`<option>${t`Loading...`}</option>`)}
                        </select>
                        <p class="pf-c-form__helper-text">${t`Stage used to configure Authenticator when user doesn't have any compatible devices. After this configuration Stage passes, the user is not prompted again.`}</p>
                    </ak-form-element-horizontal>
                    `: html``}
                </div>
            </ak-form-group>
        </form>`;
    }

}
