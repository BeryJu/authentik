import { gettext } from "django";
import { customElement, html, LitElement, property, TemplateResult } from "lit-element";
import { Application } from "../../api/Applications";
import { Provider } from "../../api/Providers";

import "../../elements/buttons/ModalButton";
import "../../elements/Spinner";

@customElement("ak-provider-related-application")
export class RelatedApplicationButton extends LitElement {

    @property({attribute: false})
    provider?: Provider;

    render(): TemplateResult {
        if (this.provider?.assigned_application_slug) {
            return html`<a href="#/applications/${this.provider.assigned_application_slug}">
                ${this.provider.assigned_application_name}
            </a>`;
        }
        return html`<ak-modal-button href=${Application.adminUrl(`create/?provider=${this.provider ? this.provider.pk : ""}`)}>
                <ak-spinner-button slot="trigger" class="pf-m-primary">
                    ${gettext("Create")}
                </ak-spinner-button>
                <div slot="modal"></div>
            </ak-modal-button>`;
    }

}
