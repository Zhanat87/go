import {Component, OnInit} from '@angular/core';
import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';

@Component({
    selector: 'home',
    styleUrls: ['./home.scss'],
    templateUrl: './home.html'
})
export class HomeComponent implements OnInit {

    constructor(private _state: GlobalState, private localStorageService: LocalStorageService) {
    }

    ngOnInit(): void {
        this.localStorageService.remove('breadCrumbs');
        this._state.notifyChanged('breadCrumbs');
    }

}