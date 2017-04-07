import { Observable } from 'rxjs/Rx';

import { ListElement } from "../entities/listElement";
import {CommonService} from "./service";

export abstract class CommonCrudService extends CommonService {

    // note: map() and mapAll() methods not must be abstract, because not all services implements crud
    // need divide common service and extends from it common crud service
    abstract map(data): any;

    abstract mapAll(data): any;

    public mapListElement(data): Observable<ListElement[]> {
        return data;
    }

    public all() {
        let params = '';

        if (this.get_params) {
            params = '?' + this.http_build_query(this.get_params);
        }

        return this.mapAll(
            this.http.get(this.getUrl() + params)
                .map(this.extractData)
                .catch(this.handleError)
        );
    }

    public create(attributes: any) {
        return this.map(
            this.http.post(this.getUrl(), attributes)
                .map(this.extractData)
                .catch(this.handleError)
        );
    }

    public find(id: number) {
        return this.map(
            this.http.get(this.getUrl() + '/' + id)
                .map(this.extractData)
                .catch(this.handleError)
        );
    }

    public update(attributes: any, id: number) {
        return this.map(
            this.http.put(this.getUrl() + '/' + id, attributes)
                .map(this.extractData)
                .catch(this.handleError)
        );
    }

    public delete(id: number) {
        return this.map(
            this.http.delete(this.getUrl() + '/' + id)
                .map(this.extractData)
                .catch(this.handleError)
        );
    }
}
