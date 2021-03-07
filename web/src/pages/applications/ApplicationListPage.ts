import { gettext } from "django";
import { customElement, html, property, TemplateResult } from "lit-element";
import { AKResponse } from "../../api/Client";
import { TablePage } from "../../elements/table/TablePage";

import "../../elements/buttons/ModalButton";
import "../../elements/buttons/SpinnerButton";
import { TableColumn } from "../../elements/table/Table";
import { PAGE_SIZE } from "../../constants";
import { Application, CoreApi } from "../../api";
import { DEFAULT_CONFIG } from "../../api/Config";
import { AdminURLManager } from "../../api/legacy";

@customElement("ak-application-list")
export class ApplicationListPage extends TablePage<Application> {
    searchEnabled(): boolean {
        return true;
    }
    pageTitle(): string {
        return gettext("Applications");
    }
    pageDescription(): string {
        return gettext("External Applications which use authentik as Identity-Provider, utilizing protocols like OAuth2 and SAML.");
    }
    pageIcon(): string {
        return gettext("pf-icon pf-icon-applications");
    }

    @property()
    order = "name";

    apiEndpoint(page: number): Promise<AKResponse<Application>> {
        return new CoreApi(DEFAULT_CONFIG).coreApplicationsList({
            ordering: this.order,
            page: page,
            pageSize: PAGE_SIZE,
            search: this.search || "",
        });
    }

    columns(): TableColumn[] {
        return [
            new TableColumn(""),
            new TableColumn("Name", "name"),
            new TableColumn("Slug", "slug"),
            new TableColumn("Provider"),
            new TableColumn("Provider Type"),
            new TableColumn(""),
        ];
    }

    row(item: Application): TemplateResult[] {
        return [
            item.metaIcon ?
                html`<img class="app-icon pf-c-avatar" src="${item.metaIcon}" alt="${gettext("Application Icon")}">` :
                html`<i class="pf-icon pf-icon-arrow"></i>`,
            html`<a href="#/core/applications/${item.slug}">
                <div>
                    ${item.name}
                </div>
                ${item.metaPublisher ? html`<small>${item.metaPublisher}</small>` : html``}
            </a>`,
            html`<code>${item.slug}</code>`,
            html`${item.provider?.name}`,
            html`${item.provider?.verboseName}`,
            html`
            <ak-modal-button href="${AdminURLManager.applications(`${item.pk}/update/`)}">
                <ak-spinner-button slot="trigger" class="pf-m-secondary">
                    ${gettext("Edit")}
                </ak-spinner-button>
                <div slot="modal"></div>
            </ak-modal-button>
            <ak-modal-button href="${AdminURLManager.applications(`${item.pk}/delete/`)}">
                <ak-spinner-button slot="trigger" class="pf-m-danger">
                    ${gettext("Delete")}
                </ak-spinner-button>
                <div slot="modal"></div>
            </ak-modal-button>
            `,
        ];
    }

    renderToolbar(): TemplateResult {
        return html`
        <ak-modal-button href=${AdminURLManager.applications("create/")}>
            <ak-spinner-button slot="trigger" class="pf-m-primary">
                ${gettext("Create")}
            </ak-spinner-button>
            <div slot="modal"></div>
        </ak-modal-button>
        ${super.renderToolbar()}
        `;
    }
}
