import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import {ArtistService} from "./artist.service";
import {BaseListPagination} from "../../common/base/baseListPagination";
import {Artist} from "./artist";
import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';

@Component({
    moduleId: 'artist',
    selector: 'artist',
    styleUrls: ['./artist-list.scss'],
    templateUrl: './artist-list.html',
})
export class ArtistList extends BaseListPagination {

    public data: Artist[];

    public listUrl = '/artists';
    public title = 'Artists';

    constructor (
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        protected localStorageService: LocalStorageService,
        public service: ArtistService) {
        super();
        console.log('test');
    }

}