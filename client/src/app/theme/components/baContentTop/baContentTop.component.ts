import {Component, OnDestroy, OnInit} from '@angular/core';

import {GlobalState} from '../../../global.state';
import {BreadCrumb} from "../../../common/entities/breadCrumb";
import {Router, NavigationEnd} from "@angular/router";
import {Subscription} from "rxjs";
import {LocalStorageService} from "angular-2-local-storage";

@Component({
    selector: 'ba-content-top',
    styleUrls: ['./baContentTop.scss'],
    templateUrl: './baContentTop.html',
})
export class BaContentTop implements OnInit, OnDestroy {

    public activePageTitle: string = 'Balu admin';

    public breadCrumbs: BreadCrumb[] = [];

    private _onRouteChange: Subscription;

    constructor(private _state: GlobalState,
                private router: Router,
                private localStorageService: LocalStorageService) {
        this._state.subscribe('menu.activeLink', (activeLink) => {
            if (activeLink) {
                this.activePageTitle = activeLink.title;
            }
            this.setBreadCrumbs();
        });
        this._state.subscribe('breadCrumbs', () => {
            this.setBreadCrumbs();
        });
    }

    ngOnInit(): void {
        this._onRouteChange = this.router.events.subscribe(event => {
            // if(event instanceof NavigationStart) {
            // }
            // NavigationEnd
            // NavigationCancel
            // NavigationError
            // RoutesRecognized
            if(event instanceof NavigationEnd) {
                this.setBreadCrumbs();
            }
        });
    }

    ngOnDestroy(): void {
        this._onRouteChange.unsubscribe();
    }

    private setBreadCrumbs(): void {
        this.breadCrumbs = [];
        if (this.router.url == '/index') {
            this.activePageTitle = 'Balu admin';
        } else {
            this.breadCrumbs.push(new BreadCrumb('Home', '/index'));
            let breadCrumbs = this.localStorageService.get('breadCrumbs') ?
                JSON.parse(this.localStorageService.get<string>('breadCrumbs')) : null;
            if (breadCrumbs) {
                this.breadCrumbs = this.breadCrumbs.concat(breadCrumbs);
            } else {
                this.breadCrumbs.push(new BreadCrumb(this.activePageTitle));
            }
        }
    }

}