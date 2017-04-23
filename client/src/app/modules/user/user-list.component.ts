import { Component, ViewChild } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import {UserService} from "./user.service";
import {BaseListPagination} from "../../common/base/baseListPagination";
import {User} from "./user";
// import {ModalDirective} from 'ngx-bootstrap';
import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';

@Component({
    moduleId: 'user',
    selector: 'user',
    styleUrls: ['./user-list.scss'],
    templateUrl: './user-list.html',
})
export class UserList extends BaseListPagination {

    // @ViewChild('bannedModal') bannedModal: ModalDirective;
    // @ViewChild('cancelBanModal') cancelBanModal: ModalDirective;

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

    // ngOnDestroy() {
    //     super.ngOnDestroy();
    //     this.localStorageService.remove('bannedUserId');
    //     this.localStorageService.remove('cancelBanUserId');
    // }

    // showBannedModal(id: number): void {
    //     this.localStorageService.set('bannedUserId');
    //     this.bannedModal.show();
    // }
    //
    // hideBannedModal(): void {
    //     this.localStorageService.remove('bannedUserId');
    //     this.bannedModal.hide();
    // }
    //
    // ban(): void {
    //     this.service.ban(this.localStorageService.get('bannedUserId'))
    //         .subscribe(
    //             data => {
    //                 if (data.success == true) {
    //                     this.getPage();
    //                     this.hideBannedModal();
    //                 } else {
    //                     console.log(data.message);
    //                 }
    //             },
    //             err => console.error(err)
    //         );
    // }
    //
    // showCancelBanModal(id: number): void {
    //     this.localStorageService.set('cancelBanUserId', id);
    //     this.cancelBanModal.show();
    // }
    //
    // hideCancelBanModal(): void {
    //     this.localStorageService.remove('cancelBanUserId');
    //     this.cancelBanModal.hide();
    // }
    //
    // cancelBan(): void {
    //     this.service.cancelBan(this.localStorageService.get('cancelBanUserId'))
    //         .subscribe(
    //             data => {
    //                 if (data.success == true) {
    //                     this.getPage();
    //                     this.hideCancelBanModal();
    //                 } else {
    //                     console.log(data.message);
    //                 }
    //             },
    //             err => console.error(err)
    //         );
    // }

}
