import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';

import {CommonListService} from "../../common/services/list.service";
import {User} from "./user";
import {Observable} from "rxjs";
import { LocalStorageService } from 'angular-2-local-storage';

@Injectable()
export class UserService extends CommonListService {

    public url = 'v1/users';

    constructor (public http: AuthHttp, protected localStorageService: LocalStorageService) {
        super();
    }

    public map(data): Observable<User> {
        return data;
    }

    public mapAll(data): Observable<User[]> {
        return data;
    }

}