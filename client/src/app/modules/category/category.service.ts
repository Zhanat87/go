import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';

import {CommonListService} from "../../common/services/list.service";
import {Category} from "./category";
import {Observable} from "rxjs";
import { LocalStorageService } from 'angular-2-local-storage';

@Injectable()
export class CategoryService extends CommonListService {

    public url = 'v1/categories';

    constructor (public http: AuthHttp, protected localStorageService: LocalStorageService) {
        super();
    }

    public map(data): Observable<Category> {
        return data;
    }

    public mapAll(data): Observable<Category[]> {
        return data;
    }

}