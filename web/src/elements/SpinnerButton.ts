import { css, customElement, html, LitElement, property } from "lit-element";
// @ts-ignore
import GlobalsStyle from "@patternfly/patternfly/base/patternfly-globals.css";
// @ts-ignore
import ButtonStyle from "@patternfly/patternfly/components/Button/button.css";
// @ts-ignore
import SpinnerStyle from "@patternfly/patternfly/components/Spinner/spinner.css";
import {
    ColorStyles,
    PRIMARY_CLASS,
    PROGRESS_CLASS,
} from "../constants";

@customElement("pb-spinner-button")
export class SpinnerButton extends LitElement {
    @property()
    isRunning = false;

    static get styles() {
        return [
            GlobalsStyle,
            ButtonStyle,
            SpinnerStyle,
            ColorStyles,
            css`
                button {
                    /* Have to use !important here, as buttons with pf-m-progress have transition already */
                    transition: all var(--pf-c-button--m-progress--TransitionDuration) ease 0s !important;
                }
            `,
        ];
    }

    constructor() {
        super();
        this.classList.add(PRIMARY_CLASS);
    }

    setLoading() {
        this.isRunning = true;
        this.classList.add(PROGRESS_CLASS);
        this.requestUpdate();
    }

    setDone(statusClass: string) {
        this.isRunning = false;
        this.classList.remove(PROGRESS_CLASS);
        this.classList.replace(PRIMARY_CLASS, statusClass);
        this.requestUpdate();
        setTimeout(() => {
            this.classList.replace(statusClass, PRIMARY_CLASS);
            this.requestUpdate();
        }, 1000);
    }

    callAction() {
        if (this.isRunning === true) {
            return;
        }
        this.setLoading();
    }

    render() {
        return html`<button
            class="pf-c-button pf-m-progress ${this.classList.toString()}"
            @click=${() => this.callAction()}
        >
            ${this.isRunning
                ? html` <span class="pf-c-button__progress">
                      <span
                          class="pf-c-spinner pf-m-md"
                          role="progressbar"
                          aria-valuetext="Loading..."
                      >
                          <span class="pf-c-spinner__clipper"></span>
                          <span class="pf-c-spinner__lead-ball"></span>
                          <span class="pf-c-spinner__tail-ball"></span>
                      </span>
                  </span>`
                : ""}
            <slot></slot>
        </button>`;
    }
}