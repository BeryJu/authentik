import { gettext } from "django";
import { customElement, html, property, TemplateResult } from "lit-element";
import { AKResponse } from "../../api/Client";
import { TablePage } from "../../elements/table/TablePage";

import { CryptoApi, CertificateKeyPair } from "../../api";

import "../../elements/buttons/ModalButton";
import "../../elements/buttons/SpinnerButton";
import { TableColumn } from "../../elements/table/Table";
import { PAGE_SIZE } from "../../constants";
import { AdminURLManager } from "../../api/legacy";
import { DEFAULT_CONFIG } from "../../api/Config";

@customElement("ak-crypto-certificatekeypair-list")
export class CertificateKeyPairListPage extends TablePage<CertificateKeyPair> {
    expandable = true;

    searchEnabled(): boolean {
        return true;
    }
    pageTitle(): string {
        return gettext("Certificate-Key Pairs");
    }
    pageDescription(): string {
        return gettext("Import certificates of external providers or create certificates to sign requests with.");
    }
    pageIcon(): string {
        return gettext("pf-icon pf-icon-key");
    }

    @property()
    order = "name";

    apiEndpoint(page: number): Promise<AKResponse<CertificateKeyPair>> {
        return new CryptoApi(DEFAULT_CONFIG).cryptoCertificatekeypairsList({
            ordering: this.order,
            page: page,
            pageSize: PAGE_SIZE,
            search: this.search || "",
        });
    }

    columns(): TableColumn[] {
        return [
            new TableColumn("Name", "name"),
            new TableColumn("Private key available?"),
            new TableColumn("Expiry date"),
            new TableColumn(""),
        ];
    }

    row(item: CertificateKeyPair): TemplateResult[] {
        return [
            html`${item.name}`,
            html`${gettext(item.privateKeyAvailable ? "Yes" : "No")}`,
            html`${item.certExpiry?.toLocaleString()}`,
            html`
            <ak-modal-button href="${AdminURLManager.cryptoCertificates(`${item.pk}/update/`)}">
                <ak-spinner-button slot="trigger" class="pf-m-secondary">
                    ${gettext("Edit")}
                </ak-spinner-button>
                <div slot="modal"></div>
            </ak-modal-button>
            <ak-modal-button href="${AdminURLManager.cryptoCertificates(`${item.pk}/delete/`)}">
                <ak-spinner-button slot="trigger" class="pf-m-danger">
                    ${gettext("Delete")}
                </ak-spinner-button>
                <div slot="modal"></div>
            </ak-modal-button>
            `,
        ];
    }

    renderExpanded(item: CertificateKeyPair): TemplateResult {
        return html`
        <td role="cell" colspan="3">
            <div class="pf-c-table__expandable-row-content">
                    <dl class="pf-c-description-list pf-m-horizontal">
                        <div class="pf-c-description-list__group">
                            <dt class="pf-c-description-list__term">
                                <span class="pf-c-description-list__text">${gettext("Certificate Fingerprint")}</span>
                            </dt>
                            <dd class="pf-c-description-list__description">
                                <div class="pf-c-description-list__text">${item.fingerprint}</div>
                            </dd>
                        </div>
                        <div class="pf-c-description-list__group">
                            <dt class="pf-c-description-list__term">
                                <span class="pf-c-description-list__text">${gettext("Certificate Subjet")}</span>
                            </dt>
                            <dd class="pf-c-description-list__description">
                                <div class="pf-c-description-list__text">${item.certSubject}</div>
                            </dd>
                        </div>
                    </dl>
            </div>
        </td>
        <td></td>
        <td></td>`;
    }

    renderToolbar(): TemplateResult {
        return html`
        <ak-modal-button href=${AdminURLManager.cryptoCertificates("create/")}>
            <ak-spinner-button slot="trigger" class="pf-m-primary">
                ${gettext("Create")}
            </ak-spinner-button>
            <div slot="modal"></div>
        </ak-modal-button>
        <ak-modal-button href=${AdminURLManager.cryptoCertificates("generate/")}>
            <ak-spinner-button slot="trigger" class="pf-m-secondary">
                ${gettext("Generate")}
            </ak-spinner-button>
            <div slot="modal"></div>
        </ak-modal-button>
        ${super.renderToolbar()}
        `;
    }
}
