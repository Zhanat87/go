import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';

import {CommonListService} from "../../common/services/list.service";
import {NewsReplication} from "./newsReplication";
import {Observable} from "rxjs";
import { LocalStorageService } from 'angular-2-local-storage';

@Injectable()
export class NewsReplicationService extends CommonListService {

    public url = 'v1/replication/news';

    constructor (public http: AuthHttp, protected localStorageService: LocalStorageService) {
        super();
    }

    public map(data): Observable<NewsReplication> {
        return data;
    }

    public mapAll(data): Observable<NewsReplication[]> {
        return data;
    }

}