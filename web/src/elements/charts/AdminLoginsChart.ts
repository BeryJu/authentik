import { customElement } from "lit-element";
import { ChartDataset } from "chart.js";
import { AdminApi, LoginMetrics } from "authentik-api";
import { AKChart } from "./Chart";
import { DEFAULT_CONFIG } from "../../api/Config";

@customElement("ak-charts-admin-login")
export class AdminLoginsChart extends AKChart<LoginMetrics> {

    apiRequest(): Promise<LoginMetrics> {
        return new AdminApi(DEFAULT_CONFIG).adminMetricsList();
    }

    getDatasets(data: LoginMetrics): ChartDataset[] {
        return [
            {
                label: "Failed Logins",
                backgroundColor: "rgba(201, 25, 11, .5)",
                spanGaps: true,
                data: data.loginsFailedPer1h?.map((cord) => {
                    return {
                        x: cord.xCord || 0,
                        y: cord.yCord || 0,
                    };
                }) || [],
            },
            {
                label: "Successful Logins",
                backgroundColor: "rgba(189, 229, 184, .5)",
                spanGaps: true,
                data: data.loginsPer1h?.map((cord) => {
                    return {
                        x: cord.xCord || 0,
                        y: cord.yCord || 0,
                    };
                }) || [],
            },
        ];
    }

}
