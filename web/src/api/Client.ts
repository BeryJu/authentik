import { NotFoundError, RequestError } from "./Error";

export const VERSION = "v2beta";

export interface QueryArguments {
    [key: string]: number | string | boolean;
}

export class Client {
    makeUrl(url: string[], query?: QueryArguments): string {
        let builtUrl = `/api/${VERSION}/${url.join("/")}/`;
        if (query) {
            const queryString = Object.keys(query)
                .map((k) => encodeURIComponent(k) + "=" + encodeURIComponent(query[k]))
                .join("&");
            builtUrl += `?${queryString}`;
        }
        return builtUrl;
    }

    fetch<T>(url: string[], query?: QueryArguments): Promise<T> {
        const finalUrl = this.makeUrl(url, query);
        return fetch(finalUrl)
            .then((r) => {
                if (r.status > 300) {
                    switch (r.status) {
                    case 404:
                        throw new NotFoundError(`URL ${finalUrl} not found`);
                    default:
                        throw new RequestError(r.statusText);
                    }
                }
                return r;
            })
            .then((r) => r.json())
            .then((r) => <T>r);
    }
}

export const DefaultClient = new Client();

export interface PBPagination {
    next?: number;
    previous?: number;

    count: number;
    current: number;
    total_pages: number;

    start_index: number;
    end_index: number;
}

export interface PBResponse<T> {
    pagination: PBPagination;

    results: Array<T>;
}
