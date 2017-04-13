import {Component, OnInit} from '@angular/core';
import {GlobalState} from "../../global.state";

@Component({
    selector: 'home',
    styleUrls: ['./home.scss'],
    templateUrl: './home.html'
})
export class Home implements OnInit {

    constructor(private _state: GlobalState) {
    }

    ngOnInit(): void {
        localStorage.removeItem('breadCrumbs');
        this._state.notifyChanged('breadCrumbs');
    }

}