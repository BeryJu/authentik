import { gettext } from "django";
import { LitElement, html, customElement, TemplateResult, property } from "lit-element";
import "./Message";
import { APIMessage } from "./Message";

export function showMessage(message: APIMessage): void {
    const container = document.querySelector<MessageContainer>("ak-message-container");
    if (!container) {
        throw new Error("failed to find message container");
    }
    container.messages.push(message);
    container.requestUpdate();
}

@customElement("ak-message-container")
export class MessageContainer extends LitElement {

    @property({attribute: false})
    messages: APIMessage[] = [];

    messageSocket?: WebSocket;
    retryDelay = 200;

    createRenderRoot(): ShadowRoot | Element {
        return this;
    }

    constructor() {
        super();
        try {
            this.connect();
        } catch (error) {
            console.warn(`authentik/messages: failed to connect to ws ${error}`);
        }
    }

    connect(): void {
        if (navigator.webdriver) return;
        const wsUrl = `${window.location.protocol.replace("http", "ws")}//${
            window.location.host
        }/ws/client/`;
        this.messageSocket = new WebSocket(wsUrl);
        this.messageSocket.addEventListener("open", () => {
            console.debug(`authentik/messages: connected to ${wsUrl}`);
            this.retryDelay = 200;
        });
        this.messageSocket.addEventListener("close", (e) => {
            console.debug(`authentik/messages: closed ws connection: ${e}`);
            if (this.retryDelay > 3000) {
                showMessage({
                    level_tag: "error",
                    message: gettext("Connection error, reconnecting...")
                });
            }
            setTimeout(() => {
                console.debug(`authentik/messages: reconnecting ws in ${this.retryDelay}ms`);
                this.connect();
            }, this.retryDelay);
            this.retryDelay = this.retryDelay * 2;
        });
        this.messageSocket.addEventListener("message", (e) => {
            const data = JSON.parse(e.data);
            this.messages.push(data);
            this.requestUpdate();
        });
        this.messageSocket.addEventListener("error", (e) => {
            console.warn(`authentik/messages: error ${e}`);
            this.retryDelay = this.retryDelay * 2;
        });
    }

    render(): TemplateResult {
        return html`<ul class="pf-c-alert-group pf-m-toast">
            ${this.messages.map((m) => {
                return html`<ak-message
                    .message=${m}
                    .onRemove=${(m: APIMessage) => {
                        this.messages = this.messages.filter((v) => v !== m);
                            this.requestUpdate();
                        }}>
                    </ak-message>`;
                })}
        </ul>`;
    }
}
