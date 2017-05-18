import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import {UserService} from "./user.service";
import {BaseListPagination} from "../../common/base/baseListPagination";
import {User} from "./user";
import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';
import {Environment} from "../../common/environment";

@Component({
    moduleId: 'user',
    selector: 'user',
    styleUrls: ['./user-list.scss'],
    templateUrl: './user-list.html',
})
export class UserList extends BaseListPagination {

    public data: User[];

    public listUrl = '/users';
    public title = 'Users';

    constructor (
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        protected localStorageService: LocalStorageService,
        public service: UserService) {
        super();
    }

    getUserPhone(user: User): string {
        if (user.phones) {
            let phones = JSON.parse(user.phones);
            return phones.phone;
        }
        return '';
    }

    getUserMobilePhone(user: User): string {
        if (user.phones) {
            let phones = JSON.parse(user.phones);
            return phones.mobile_phone;
        }
        return '';
    }

    getAvatar(user: User): string {
        return user.avatar_string ? (user.avatar_string.substr(0, 4) == 'http' ? user.avatar_string :
            Environment.API_ENDPOINT + 'static/users/avatars/' + user.avatar_string) : '';
    }

}
