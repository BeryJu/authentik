import { css, CSSResult, customElement, html, LitElement, property, TemplateResult } from "lit-element";
import Chart from "chart.js";
import { LoginMetrics } from "../api";

interface TickValue {
    value: number;
    major: boolean;
}

@customElement("ak-admin-logins-chart")
export class AdminLoginsChart extends LitElement {
    @property({type: Array})
    url: string[] = [];

    chart?: Chart;

    static get styles(): CSSResult[] {
        return [css`
            :host {
                position: relative;
                height: 100%;
                width: 100%;
                display: block;
                min-height: 25rem;
            }
            canvas {
                width: 100px;
                height: 100px;
            }
        `];
    }

    constructor() {
        super();
        window.addEventListener("resize", () => {
            if (this.chart) {
                this.chart.resize();
            }
        });
    }

    @property({ attribute: false })
    apiRequest(): Promise<LoginMetrics> {
        throw new Error();
    }

    firstUpdated(): void {
        this.apiRequest().then((r) => {
            const canvas = <HTMLCanvasElement>this.shadowRoot?.querySelector("canvas");
            if (!canvas) {
                console.warn("Failed to get canvas element");
                return false;
            }
            const ctx = canvas.getContext("2d");
            if (!ctx) {
                console.warn("failed to get 2d context");
                return false;
            }
            this.chart = new Chart(ctx, {
                type: "bar",
                data: {
                    datasets: [
                        {
                            label: "Failed Logins",
                            backgroundColor: "rgba(201, 25, 11, .5)",
                            spanGaps: true,
                            data: r.loginsFailedPer1h?.map((cord) => {
                                return {
                                    x: cord.xCord,
                                    y: cord.yCord,
                                };
                            }),
                        },
                        {
                            label: "Successful Logins",
                            backgroundColor: "rgba(189, 229, 184, .5)",
                            spanGaps: true,
                            data: r.loginsPer1h?.map((cord) => {
                                return {
                                    x: cord.xCord,
                                    y: cord.yCord,
                                };
                            }),
                        },
                    ],
                },
                options: {
                    maintainAspectRatio: false,
                    spanGaps: true,
                    scales: {
                        xAxes: [
                            {
                                stacked: true,
                                gridLines: {
                                    color: "rgba(0, 0, 0, 0)",
                                },
                                type: "time",
                                offset: true,
                                ticks: {
                                    callback: function (value, index: number, values) {
                                        const valueStamp = <TickValue>(<unknown>values[index]);
                                        const delta = Date.now() - valueStamp.value;
                                        const ago = Math.round(delta / 1000 / 3600);
                                        return `${ago} Hours ago`;
                                    },
                                    autoSkip: true,
                                    maxTicksLimit: 8,
                                },
                            },
                        ],
                        yAxes: [
                            {
                                stacked: true,
                                gridLines: {
                                    color: "rgba(0, 0, 0, 0)",
                                },
                            },
                        ],
                    },
                },
            });
        });
    }

    render(): TemplateResult {
        return html`<canvas></canvas>`;
    }
}
