import { gettext } from "django";
import { customElement, html, property, TemplateResult } from "lit-element";
import { AKResponse } from "../../api/Client";
import { Table, TableColumn } from "../../elements/table/Table";

import "../../elements/Tabs";
import "../../elements/buttons/ModalButton";
import "../../elements/buttons/SpinnerButton";
import "../../elements/buttons/Dropdown";
import "../../elements/policies/BoundPoliciesList";
import { until } from "lit-html/directives/until";
import { PAGE_SIZE } from "../../constants";
import { FlowsApi, FlowStageBinding, StagesApi } from "../../api";
import { DEFAULT_CONFIG } from "../../api/Config";
import { AdminURLManager } from "../../api/legacy";

@customElement("ak-bound-stages-list")
export class BoundStagesList extends Table<FlowStageBinding> {
    expandable = true;

    @property()
    target?: string;

    apiEndpoint(page: number): Promise<AKResponse<FlowStageBinding>> {
        return new FlowsApi(DEFAULT_CONFIG).flowsBindingsList({
            target: this.target || "",
            ordering: "order",
            page: page,
            pageSize: PAGE_SIZE,
        });
    }

    columns(): TableColumn[] {
        return [
            new TableColumn("Order"),
            new TableColumn("Name"),
            new TableColumn("Type"),
            new TableColumn(""),
        ];
    }

    row(item: FlowStageBinding): TemplateResult[] {
        return [
            html`${item.order}`,
            html`${item.stageObj?.name}`,
            html`${item.stageObj?.verboseName}`,
            html`
            <ak-modal-button href="${AdminURLManager.stageBindings(`${item.pk}/update/`)}">
                <ak-spinner-button slot="trigger" class="pf-m-secondary">
                    ${gettext("Edit Binding")}
                </ak-spinner-button>
                <div slot="modal"></div>
            </ak-modal-button>
            <ak-modal-button href="${AdminURLManager.stages(`${item.stage}/update/`)}">
                <ak-spinner-button slot="trigger" class="pf-m-secondary">
                    ${gettext("Edit Stage")}
                </ak-spinner-button>
                <div slot="modal"></div>
            </ak-modal-button>
            <ak-modal-button href="${AdminURLManager.stages(`${item.pk}/delete/`)}">
                <ak-spinner-button slot="trigger" class="pf-m-danger">
                    ${gettext("Delete")}
                </ak-spinner-button>
                <div slot="modal"></div>
            </ak-modal-button>
            `,
        ];
    }

    renderExpanded(item: FlowStageBinding): TemplateResult {
        return html`
        <td></td>
        <td role="cell" colspan="3">
            <div class="pf-c-table__expandable-row-content">
                <div class="pf-c-content">
                    <p>${gettext("These policies control when this stage will be applied to the flow.")}</p>
                    <ak-bound-policies-list .target=${item.policybindingmodelPtrId}>
                    </ak-bound-policies-list>
                </div>
            </div>
        </td>
        <td></td>
        <td></td>`;
    }

    renderEmpty(): TemplateResult {
        return super.renderEmpty(html`<ak-empty-state header=${gettext("No Stages bound")} icon="pf-icon-module">
            <div slot="body">
                ${gettext("No stages are currently bound to this flow.")}
            </div>
            <div slot="primary">
                <ak-modal-button href="${AdminURLManager.stageBindings(`create/?target=${this.target}`)}">
                    <ak-spinner-button slot="trigger" class="pf-m-primary">
                        ${gettext("Bind Stage")}
                    </ak-spinner-button>
                    <div slot="modal"></div>
                </ak-modal-button>
            </div>
        </ak-empty-state>`);
    }

    renderToolbar(): TemplateResult {
        return html`
        <ak-dropdown class="pf-c-dropdown">
            <button class="pf-m-primary pf-c-dropdown__toggle" type="button">
                <span class="pf-c-dropdown__toggle-text">${gettext("Create Stage")}</span>
                <i class="fas fa-caret-down pf-c-dropdown__toggle-icon" aria-hidden="true"></i>
            </button>
            <ul class="pf-c-dropdown__menu" hidden>
                ${until(new StagesApi(DEFAULT_CONFIG).stagesAllTypes({}).then((types) => {
                    return types.map((type) => {
                        return html`<li>
                            <ak-modal-button href="${type.link}">
                                <button slot="trigger" class="pf-c-dropdown__menu-item">${type.name}<br>
                                    <small>${type.description}</small>
                                </button>
                                <div slot="modal"></div>
                            </ak-modal-button>
                        </li>`;
                    });
                }), html`<ak-spinner></ak-spinner>`)}
            </ul>
        </ak-dropdown>
        <ak-modal-button href="${AdminURLManager.stageBindings(`create/?target=${this.target}`)}">
            <ak-spinner-button slot="trigger" class="pf-m-primary">
                ${gettext("Bind Stage")}
            </ak-spinner-button>
            <div slot="modal"></div>
        </ak-modal-button>
        ${super.renderToolbar()}
        `;
    }
}
