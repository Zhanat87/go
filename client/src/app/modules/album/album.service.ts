import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';

import {CommonListService} from "../../common/services/list.service";
import {Album} from "./album";
import {Observable} from "rxjs";

@Injectable()
export class AlbumService extends CommonListService {

    public url = 'v1/albums';

    constructor (public http: AuthHttp) {
        super();
    }

    public map(data): Observable<Album> {
        return data;
    }

    public mapAll(data): Observable<Album[]> {
        return data;
    }

}