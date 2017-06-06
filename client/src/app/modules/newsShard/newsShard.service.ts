import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';

import {CommonListService} from "../../common/services/list.service";
import {NewsShard} from "./newsShard";
import {Observable} from "rxjs";
import { LocalStorageService } from 'angular-2-local-storage';

@Injectable()
export class NewsShardService extends CommonListService {

    public url = 'v1/shard/news';

    constructor (public http: AuthHttp, protected localStorageService: LocalStorageService) {
        super();
    }

    public map(data): Observable<NewsShard> {
        return data;
    }

    public mapAll(data): Observable<NewsShard[]> {
        return data;
    }

}