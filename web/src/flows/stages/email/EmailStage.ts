import { gettext } from "django";
import { CSSResult, customElement, html, property, TemplateResult } from "lit-element";
import { Challenge } from "../../../api";
import { COMMON_STYLES } from "../../../common/styles";
import { BaseStage } from "../base";
import "../../../elements/utils/LoadingState";

export type EmailChallenge = Challenge;

@customElement("ak-stage-email")
export class EmailStage extends BaseStage {

    @property({ attribute: false })
    challenge?: EmailChallenge;

    static get styles(): CSSResult[] {
        return COMMON_STYLES;
    }

    render(): TemplateResult {
        if (!this.challenge) {
            return html`<ak-loading-state></ak-loading-state>`;
        }
        return html`<header class="pf-c-login__main-header">
                <h1 class="pf-c-title pf-m-3xl">
                    ${this.challenge.title}
                </h1>
            </header>
            <div class="pf-c-login__main-body">
                <form class="pf-c-form" @submit=${(e: Event) => { this.submitForm(e); }}>
                    <div class="pf-c-form__group">
                        <p>
                            ${gettext("Check your Emails for a password reset link.")}
                        </p>
                    </div>

                    <div class="pf-c-form__group pf-m-action">
                        <button type="submit" class="pf-c-button pf-m-primary pf-m-block">
                            ${gettext("Send Email again.")}
                        </button>
                    </div>
                </form>
            </div>
            <footer class="pf-c-login__main-footer">
                <ul class="pf-c-login__main-footer-links">
                </ul>
            </footer>`;
    }

}
