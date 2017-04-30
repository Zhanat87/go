import {CommonCrudService} from "./crud.service";

export abstract class CommonListService extends CommonCrudService {

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

    public allWithoutLimit() {
        return this.mapAll(
            this.http.get(this.getUrl() + '?limit=false')
                .map(this.extractItems)
                .catch(this.handleError)
        );
    }

    public create(attributes: any) {
        return this.map(
            this.http.post(this.getUrl(), attributes)
                .map(this.extractAllData)
                .catch(this.handleServerErrors)
        );
    }

    public find(id: number) {
        return this.map(
            this.http.get(this.getUrl() + '/' + id)
                .map(this.extractAllData)
                .catch(this.handleError)
        );
    }

    public update(attributes: any, id: number) {
        return this.map(
            this.http.put(this.getUrl() + '/' + id, attributes)
                .map(this.extractAllData)
                .catch(this.handleServerErrors)
        );
    }

    public delete(id: number) {
        return this.map(
            this.http.delete(this.getUrl() + '/' + id)
                .map(this.extractAllData)
                .catch(this.handleError)
        );
    }

    public ban(id: number) {
        return this.update({'banned': true}, id);
    }

    public cancelBan(id: number) {
        return this.update({'banned': false}, id);
    }

    public publish(id: number) {
        return this.update({'status_id': 'ACTIVE'}, id);
    }

    public unPublish(id: number) {
        return this.update({'status_id': 'PENDING'}, id);
    }

}
