import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';

import {CommonListService} from "../../common/services/list.service";
import {Album} from "./album";
import {Observable} from "rxjs";
import { LocalStorageService } from 'angular-2-local-storage';

@Injectable()
export class AlbumService extends CommonListService {

    public url = 'v1/albums';

    constructor (public http: AuthHttp, protected localStorageService: LocalStorageService) {
        super();
    }

    public map(data): Observable<Album> {
        return data;
    }

    public mapAll(data): Observable<Album[]> {
        return data;
    }

}