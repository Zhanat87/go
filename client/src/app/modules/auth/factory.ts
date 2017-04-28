import { Http, RequestOptions } from '@angular/http';
import { AuthHttp, AuthConfig } from 'angular2-jwt';
import {trim} from "../../common/utils";

export function authHttpServiceFactory(http: Http, options: RequestOptions) {
    return new AuthHttp(new AuthConfig({
        tokenName: 'token',
        tokenGetter: (() => trim(localStorage.getItem('id_token'), '"')),
        globalHeaders: [{'Content-Type':'application/json'}],
    }), http, options);
}