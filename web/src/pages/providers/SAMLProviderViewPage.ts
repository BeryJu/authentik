import { gettext } from "django";
import { CSSResult, customElement, html, property, TemplateResult } from "lit-element";
import { until } from "lit-html/directives/until";
import { COMMON_STYLES } from "../../common/styles";

import "../../elements/buttons/ModalButton";
import "../../elements/buttons/SpinnerButton";
import "../../elements/CodeMirror";
import "../../elements/Tabs";
import { Page } from "../../elements/Page";
import "./RelatedApplicationButton";
import { ProvidersApi, SAMLProvider } from "../../api";
import { DEFAULT_CONFIG } from "../../api/Config";
import { AdminURLManager, AppURLManager } from "../../api/legacy";

@customElement("ak-provider-saml-view")
export class SAMLProviderViewPage extends Page {
    pageTitle(): string {
        return gettext(`SAML Provider ${this.provider?.name || ""}`);
    }
    pageDescription(): string | undefined {
        return;
    }
    pageIcon(): string {
        return "pf-icon pf-icon-integration";
    }

    @property()
    set args(value: { [key: string]: number }) {
        this.providerID = value.id;
    }

    @property({type: Number})
    set providerID(value: number) {
        new ProvidersApi(DEFAULT_CONFIG).providersSamlRead({
            id: value,
        }).then((prov) => (this.provider = prov));
    }

    @property({ attribute: false })
    provider?: SAMLProvider;

    static get styles(): CSSResult[] {
        return COMMON_STYLES;
    }

    constructor() {
        super();
        this.addEventListener("ak-refresh", () => {
            if (!this.provider?.pk) return;
            this.providerID = this.provider?.pk;
        });
    }

    renderContent(): TemplateResult {
        if (!this.provider) {
            return html``;
        }
        return html`<ak-tabs>
                <section slot="page-1" data-tab-title="${gettext("Overview")}" class="pf-c-page__main-section pf-m-no-padding-mobile">
                    <div class="pf-u-display-flex pf-u-justify-content-center">
                        <div class="pf-u-w-75">
                            <div class="pf-c-card pf-c-card-aggregate">
                                <div class="pf-c-card__body">
                                    <dl class="pf-c-description-list pf-m-3-col-on-lg">
                                        <div class="pf-c-description-list__group">
                                            <dt class="pf-c-description-list__term">
                                                <span class="pf-c-description-list__text">${gettext("Name")}</span>
                                            </dt>
                                            <dd class="pf-c-description-list__description">
                                                <div class="pf-c-description-list__text">${this.provider.name}</div>
                                            </dd>
                                        </div>
                                        <div class="pf-c-description-list__group">
                                            <dt class="pf-c-description-list__term">
                                                <span class="pf-c-description-list__text">${gettext("Assigned to application")}</span>
                                            </dt>
                                            <dd class="pf-c-description-list__description">
                                                <div class="pf-c-description-list__text">
                                                    <ak-provider-related-application .provider=${this.provider}></ak-provider-related-application>
                                                </div>
                                            </dd>
                                        </div>
                                        <div class="pf-c-description-list__group">
                                            <dt class="pf-c-description-list__term">
                                                <span class="pf-c-description-list__text">${gettext("ACS URL")}</span>
                                            </dt>
                                            <dd class="pf-c-description-list__description">
                                                <div class="pf-c-description-list__text">${this.provider.acsUrl}</div>
                                            </dd>
                                        </div>
                                        <div class="pf-c-description-list__group">
                                            <dt class="pf-c-description-list__term">
                                                <span class="pf-c-description-list__text">${gettext("Audience")}</span>
                                            </dt>
                                            <dd class="pf-c-description-list__description">
                                                <div class="pf-c-description-list__text">${this.provider.audience}</div>
                                            </dd>
                                        </div>
                                        <div class="pf-c-description-list__group">
                                            <dt class="pf-c-description-list__term">
                                                <span class="pf-c-description-list__text">${gettext("Issuer")}</span>
                                            </dt>
                                            <dd class="pf-c-description-list__description">
                                                <div class="pf-c-description-list__text">${this.provider.issuer}</div>
                                            </dd>
                                        </div>
                                    </dl>
                                </div>
                                <div class="pf-c-card__footer">
                                    <ak-modal-button href="${AdminURLManager.providers(`${this.provider.pk}/update/`)}">
                                        <ak-spinner-button slot="trigger" class="pf-m-primary">
                                            ${gettext("Edit")}
                                        </ak-spinner-button>
                                        <div slot="modal"></div>
                                    </ak-modal-button>
                                </div>
                            </div>
                        </div>
                    </div>
                </section>
                <section slot="page-2" data-tab-title="${gettext("Metadata")}" class="pf-c-page__main-section pf-m-no-padding-mobile">
                    <div class="pf-u-display-flex pf-u-justify-content-center">
                        <div class="pf-u-w-75">
                            <div class="pf-c-card pf-c-card-aggregate">
                                <div class="pf-c-card__body">
                                    ${until(
                                        new ProvidersApi(DEFAULT_CONFIG).providersSamlMetadata({
                                            id: this.provider.pk || 0,
                                        }).then(m => {
                                            return html`<ak-codemirror mode="xml"><textarea class="pf-c-form-control" readonly>${m.metadata}</textarea></ak-codemirror>`;
                                        })
                                    )}
                                </div>
                                <div class="pf-c-card__footer">
                                    <a class="pf-c-button pf-m-primary" target="_blank" href="${AppURLManager.providerSAML(`${this.provider.assignedApplicationSlug}/metadata/`)}">
                                        ${gettext("Download")}
                                    </a>
                                </div>
                            </div>
                        </div>
                    </div>
                </section>
            </ak-tabs>`;
    }
}
