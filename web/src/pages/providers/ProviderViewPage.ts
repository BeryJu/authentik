import { CSSResult, customElement, html, LitElement, property, TemplateResult } from "lit-element";
import { Provider } from "../../api/Providers";
import { COMMON_STYLES } from "../../common/styles";

import "../../elements/buttons/ModalButton";
import "../../elements/buttons/SpinnerButton";
import "../../elements/utils/LoadingState";

import "./SAMLProviderViewPage";
import "./OAuth2ProviderViewPage";
import "./ProxyProviderViewPage";

@customElement("ak-provider-view")
export class ProviderViewPage extends LitElement {

    @property({type: Number})
    set providerID(value: number) {
        Provider.get(value).then((app) => (this.provider = app));
    }

    @property({ attribute: false })
    provider?: Provider;

    static get styles(): CSSResult[] {
        return COMMON_STYLES;
    }

    render(): TemplateResult {
        if (!this.provider) {
            return html`<ak-loading-state></ak-loading-state>`;
        }
        switch (this.provider?.object_type) {
            case "saml":
                return html`<ak-provider-saml-view providerID=${this.provider.pk}></ak-provider-saml-view>`;
            case "oauth2":
                return html`<ak-provider-oauth2-view providerID=${this.provider.pk}></ak-provider-oauth2-view>`;
            case "proxy":
                return html`<ak-provider-proxy-view providerID=${this.provider.pk}></ak-provider-proxy-view>`;
            default:
                return html`<p>Invalid provider type ${this.provider?.object_type}</p>`;
        }
    }
}
