import { EmailStage, StagesApi } from "authentik-api";
import { t } from "@lingui/macro";
import { customElement, property } from "lit-element";
import { html, TemplateResult } from "lit-html";
import { DEFAULT_CONFIG } from "../../../api/Config";
import { Form } from "../../../elements/forms/Form";
import { ifDefined } from "lit-html/directives/if-defined";
import "../../../elements/forms/HorizontalFormElement";
import "../../../elements/forms/FormGroup";
import { first } from "../../../utils";
import { until } from "lit-html/directives/until";

@customElement("ak-stage-email-form")
export class EmailStageForm extends Form<EmailStage> {

    set stageUUID(value: string) {
        new StagesApi(DEFAULT_CONFIG).stagesEmailRead({
            stageUuid: value,
        }).then(stage => {
            this.stage = stage;
        });
    }

    @property({attribute: false})
    stage?: EmailStage;

    @property({type: Boolean})
    showConnectionSettings = false;

    getSuccessMessage(): string {
        if (this.stage) {
            return t`Successfully updated stage.`;
        } else {
            return t`Successfully created stage.`;
        }
    }

    send = (data: EmailStage): Promise<EmailStage> => {
        if (this.stage) {
            return new StagesApi(DEFAULT_CONFIG).stagesEmailPartialUpdate({
                stageUuid: this.stage.pk || "",
                data: data
            });
        } else {
            return new StagesApi(DEFAULT_CONFIG).stagesEmailCreate({
                data: data
            });
        }
    };

    renderConnectionSettings(): TemplateResult {
        if (!this.showConnectionSettings) {
            return html``;
        }
        return html`<ak-form-group>
                <span slot="header">
                    ${t`Connection settings`}
                </span>
                <div slot="body" class="pf-c-form">
                    <ak-form-element-horizontal
                        label=${t`SMTP Host`}
                        ?required=${true}
                        name="host">
                        <input type="text" value="${ifDefined(this.stage?.host || "")}" class="pf-c-form-control" required>
                    </ak-form-element-horizontal>
                    <ak-form-element-horizontal
                        label=${t`SMTP Port`}
                        ?required=${true}
                        name="port">
                        <input type="number" value="${first(this.stage?.port, 25)}" class="pf-c-form-control" required>
                    </ak-form-element-horizontal>
                    <ak-form-element-horizontal
                        label=${t`SMTP Username`}
                        name="username">
                        <input type="text" value="${ifDefined(this.stage?.username || "")}" class="pf-c-form-control">
                    </ak-form-element-horizontal>
                    <ak-form-element-horizontal
                        label=${t`SMTP Password`}
                        ?writeOnly=${this.stage !== undefined}
                        name="password">
                        <input type="text" value="${ifDefined(this.stage?.password || "")}" class="pf-c-form-control">
                    </ak-form-element-horizontal>
                    <ak-form-element-horizontal name="useTls">
                        <div class="pf-c-check">
                            <input type="checkbox" class="pf-c-check__input" ?checked=${first(this.stage?.useTls, true)}>
                            <label class="pf-c-check__label">
                                ${t`Use TLS`}
                            </label>
                        </div>
                    </ak-form-element-horizontal>
                    <ak-form-element-horizontal name="useSsl">
                        <div class="pf-c-check">
                            <input type="checkbox" class="pf-c-check__input" ?checked=${first(this.stage?.useSsl, false)}>
                            <label class="pf-c-check__label">
                                ${t`Use SSL`}
                            </label>
                        </div>
                    </ak-form-element-horizontal>
                    <ak-form-element-horizontal
                        label=${t`Timeout`}
                        ?required=${true}
                        name="timeout">
                        <input type="number" value="${first(this.stage?.timeout, 30)}" class="pf-c-form-control" required>
                    </ak-form-element-horizontal>
                    <ak-form-element-horizontal
                        label=${t`From address`}
                        ?required=${true}
                        name="fromAddress">
                        <input type="text" value="${ifDefined(this.stage?.fromAddress || "system@authentik.local")}" class="pf-c-form-control" required>
                    </ak-form-element-horizontal>
                </div>
            </ak-form-group>`;
    }

    renderForm(): TemplateResult {
        return html`<form class="pf-c-form pf-m-horizontal">
            <div class="form-help-text">
                ${t`Verify the user's email address by sending them a one-time-link. Can also be used for recovery to verify the user's authenticity.`}
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
                    <ak-form-element-horizontal name="useGlobalSettings">
                        <div class="pf-c-check">
                            <input type="checkbox" class="pf-c-check__input" ?checked=${first(this.stage?.useGlobalSettings, true)} @change=${(ev: Event) => {
                                const target = ev.target as HTMLInputElement;
                                this.showConnectionSettings = !target.checked;
                            }}>
                            <label class="pf-c-check__label">
                                ${t`Use global settings`}
                            </label>
                        </div>
                        <p class="pf-c-form__helper-text">${t`When enabled, global Email connection settings will be used and connection settings below will be ignored.`}</p>
                    </ak-form-element-horizontal>
                    <ak-form-element-horizontal
                        label=${t`Token expiry`}
                        ?required=${true}
                        name="tokenExpiry">
                        <input type="number" value="${first(this.stage?.tokenExpiry, 30)}" class="pf-c-form-control" required>
                        <p class="pf-c-form__helper-text">${t`Time in minutes the token sent is valid.`}</p>
                    </ak-form-element-horizontal>
                    <ak-form-element-horizontal
                        label=${t`Subject`}
                        ?required=${true}
                        name="subject">
                        <input type="text" value="${first(this.stage?.subject, "authentik")}" class="pf-c-form-control" required>
                    </ak-form-element-horizontal>
                    <ak-form-element-horizontal
                        label=${t`Template`}
                        ?required=${true}
                        name="template">
                        <select name="users" class="pf-c-form-control">
                            ${until(new StagesApi(DEFAULT_CONFIG).stagesEmailTemplates().then(templates => {
                                return templates.map(template => {
                                    const selected = this.stage?.template === template.name;
                                    return html`<option value=${ifDefined(template.name)} ?selected=${selected}>
                                        ${template.description}
                                    </option>`;
                                });
                            }), html`<option>${t`Loading...`}</option>`)}
                        </select>
                    </ak-form-element-horizontal>
                </div>
            </ak-form-group>
            ${this.renderConnectionSettings()}
        </form>`;
    }

}
