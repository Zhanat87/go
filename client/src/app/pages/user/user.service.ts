import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';

import {CommonListService} from "../../common/services/list.service";
import {User} from "./user";
import {Observable} from "rxjs";

@Injectable()
export class UserService extends CommonListService {

    public url = 'crud/users';

    constructor (public http: AuthHttp) {
        super();
    }

    public map(data): Observable<User> {
        return data;
    }

    public mapAll(data): Observable<User[]> {
        return data;
    }

}