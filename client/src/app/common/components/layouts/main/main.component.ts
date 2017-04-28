import {Component, OnInit} from '@angular/core';
import {Routes} from '@angular/router';

import {BaMenuService} from "../../../../theme/services/baMenu/baMenu.service";
import {APP_MENU} from "../../../../app.menu";
import {Environment} from "../../../environment";

@Component({
    selector: 'main-layout',
    templateUrl: './main.html',
})
export class MainLayoutComponent implements OnInit {

    public isLocal: boolean;

    constructor(private _menuService: BaMenuService,) {
    }

    ngOnInit(): void {
        this._menuService.updateMenuByRoutes(<Routes>APP_MENU);

        this.isLocal = Environment.IS_LOCAL;
    }

}