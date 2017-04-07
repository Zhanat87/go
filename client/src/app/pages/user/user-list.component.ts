import { Component, ViewChild } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import {UserService} from "./user.service";
import {BaseListPagination} from "../../common/base/baseListPagination";
import {User} from "./user";
import {ModalDirective} from 'ngx-bootstrap';
import {GlobalState} from "../../global.state";

@Component({
    moduleId: 'user',
    selector: 'user',
    styleUrls: ['./user-list.scss'],
    templateUrl: './user-list.html',
})
export class UserList extends BaseListPagination {

    @ViewChild('bannedModal') bannedModal: ModalDirective;
    @ViewChild('cancelBanModal') cancelBanModal: ModalDirective;

    public data: User[];

    public listUrl = '/pages/users';
    public title = 'Users';

    constructor (
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        public service: UserService) {
        super();
    }

    ngOnDestroy() {
        super.ngOnDestroy();
        localStorage.removeItem('bannedUserId');
        localStorage.removeItem('cancelBanUserId');
    }

    showBannedModal(id: number): void {
        localStorage.setItem('bannedUserId', id.toString());
        this.bannedModal.show();
    }

    hideBannedModal(): void {
        localStorage.removeItem('bannedUserId');
        this.bannedModal.hide();
    }

    ban(): void {
        this.service.ban(parseInt(localStorage.getItem('bannedUserId')))
            .subscribe(
                data => {
                    if (data.success == true) {
                        this.getPage();
                        this.hideBannedModal();
                    } else {
                        console.log(data.message);
                    }
                },
                err => console.error(err)
            );
    }

    showCancelBanModal(id: number): void {
        localStorage.setItem('cancelBanUserId', id.toString());
        this.cancelBanModal.show();
    }

    hideCancelBanModal(): void {
        localStorage.removeItem('cancelBanUserId');
        this.cancelBanModal.hide();
    }

    cancelBan(): void {
        this.service.cancelBan(parseInt(localStorage.getItem('cancelBanUserId')))
            .subscribe(
                data => {
                    if (data.success == true) {
                        this.getPage();
                        this.hideCancelBanModal();
                    } else {
                        console.log(data.message);
                    }
                },
                err => console.error(err)
            );
    }

}