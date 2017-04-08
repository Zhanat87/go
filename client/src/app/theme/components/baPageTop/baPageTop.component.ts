import {Component} from '@angular/core';

import {GlobalState} from '../../../global.state';

import 'style-loader!./baPageTop.scss';
import {User} from "../../../pages/user/user";
import {LogoutService} from '../../../pages/login/logout.service';
import {LogoutResponse} from "../../../pages/login/logout.response";
import {Router} from "@angular/router";
import {tokenNotExpired} from "angular2-jwt";

@Component({
    selector: 'ba-page-top',
    templateUrl: './baPageTop.html',
})
export class BaPageTop {

    public isScrolled: boolean = false;
    public isMenuCollapsed: boolean = false;
    public user: User;

    constructor(private _state: GlobalState,
                private router: Router,
                private logoutService: LogoutService) {
        this._state.subscribe('menu.isCollapsed', (isCollapsed) => {
            this.isMenuCollapsed = isCollapsed;
        });

        this.user = JSON.parse(localStorage.getItem('currentUser')) as User;
    }

    public toggleMenu() {
        this.isMenuCollapsed = !this.isMenuCollapsed;
        this._state.notifyDataChanged('menu.isCollapsed', this.isMenuCollapsed);
        return false;
    }

    public scrolledChanged(isScrolled) {
        this.isScrolled = isScrolled;
    }

    public signOut(): void {
        if (!tokenNotExpired(null, localStorage.getItem('id_token'))) {
            console.log('token expired');
            this.cleanAndQuit();
        } else {
            this.logoutService.signOut()
                .subscribe(
                    data => {
                        let response = data as LogoutResponse;

                        if (response.message == 'token_invalidated') {
                            this.cleanAndQuit();
                        } else  {
                            console.log('error');
                        }
                    },
                    error => {
                        console.log(error);
                    },
                );
        }
    }

    redirectToLogin(): void {
        this.router.navigate(['/login']);
    }

    cleanAndQuit(): void {
        this.user = null;
        localStorage.clear();
        this.redirectToLogin();
    }

}