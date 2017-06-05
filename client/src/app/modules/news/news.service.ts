import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';

import {CommonListService} from "../../common/services/list.service";
import {News} from "./news";
import {Observable} from "rxjs";
import { LocalStorageService } from 'angular-2-local-storage';

@Injectable()
export class NewsService extends CommonListService {

    public url = 'v1/news';

    constructor (public http: AuthHttp, protected localStorageService: LocalStorageService) {
        super();
    }

    public map(data): Observable<News> {
        return data;
    }

    public mapAll(data): Observable<News[]> {
        return data;
    }

}