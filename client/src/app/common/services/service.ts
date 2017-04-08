import { Response, URLSearchParams} from '@angular/http';
import { Observable } from 'rxjs/Rx';
import {Environment} from "../environment";
import {LocalStorageService} from "angular-2-local-storage";

export abstract class CommonService {

    public http;

    public url;

    public params = {};

    public get_params = {};

    protected localStorageService: LocalStorageService;

    public getUrl() {
        let url = this.url;
        return this.parseUrl(url);
    }
    
    public parseUrl(url) {
        for(let key in this.params) {
            url = url.replace(':' + key, this.params[key]);
        }

        // return url;
        return Environment.API_ENDPOINT + url;
    }

    public extractData(res: Response) {
        let body = res.json();
        return body.data || { };
    }

    public extractAllData(res: Response) {
        let body = res.json();
        return body || { };
    }

    public handleError (error: any) {
        let errMsg = (error.message) ? error.message :
            error.status ? `${error.status} - ${error.statusText}` : 'Server error';

        console.error(errMsg); // log to console instead
        return Observable.throw(errMsg);
    }

    public http_build_query(params) {
        let urlSearchParams = new URLSearchParams();

        for (let key in params) {
            urlSearchParams.append(key, params[key]);
        }

        return urlSearchParams.toString();
    }

    public paginate(params) {
        return this.http.get(this.getUrl() + '?page=' + params.page + '&search=' + params.search + '&per_page=' + params.perPage);
    }

}