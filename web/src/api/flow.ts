import { DefaultClient, PBResponse, QueryArguments } from "./client";

export enum FlowDesignation {
    Authentication = "authentication",
    Authorization = "authorization",
    Invalidation = "invalidation",
    Enrollment = "enrollment",
    Unrenollment = "unenrollment",
    Recovery = "recovery",
    StageConfiguration = "stage_configuration",
}

export class Flow {
    pk: string;
    policybindingmodel_ptr_id: string;
    name: string;
    slug: string;
    title: string;
    designation: FlowDesignation;
    background: string;
    stages: string[];
    policies: string[];
    cache_count: number;

    constructor() {
        throw Error();
    }

    static get(slug: string): Promise<Flow> {
        return DefaultClient.fetch<Flow>(["flows", "instances", slug]);
    }

    static list(filter?: QueryArguments): Promise<PBResponse<Flow>> {
        return DefaultClient.fetch<PBResponse<Flow>>(["flows", "instances"], filter);
    }
}

export class FlowStageBinding {

    pk: string;
    target: string;
    stage: string;
    evaluate_on_plan: boolean;
    re_evaluate_policies: boolean;
    order: number;
    policies: string[];

    constructor() {
        throw Error();
    }

    static get(slug: string): Promise<FlowStageBinding> {
        return DefaultClient.fetch<FlowStageBinding>(["flows", "bindings", slug]);
    }

    static list(filter?: QueryArguments): Promise<PBResponse<FlowStageBinding>> {
        return DefaultClient.fetch<PBResponse<FlowStageBinding>>(["flows", "bindings"], filter);
    }
}
