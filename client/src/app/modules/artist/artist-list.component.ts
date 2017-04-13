import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import {ArtistService} from "./artist.service";
import {BaseListPagination} from "../../common/base/baseListPagination";
import {Artist} from "./artist";
import {GlobalState} from "../../global.state";

@Component({
    moduleId: 'artist',
    selector: 'artist',
    styleUrls: ['./artist-list.scss'],
    templateUrl: './artist-list.html',
})
export class ArtistList extends BaseListPagination {

    public data: Artist[];

    public listUrl = '/pages/artists';
    public title = 'Artists';

    constructor (
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        public service: ArtistService) {
        super();
    }

}