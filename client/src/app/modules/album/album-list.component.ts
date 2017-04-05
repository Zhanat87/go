import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import {AlbumService} from "./album.service";
import {BaseListPagination} from "../../common/base/baseListPagination";
import {Album} from "./album";
import {GlobalState} from "../../global.state";

@Component({
    moduleId: 'album',
    selector: 'album',
    styleUrls: ['./album-list.scss'],
    templateUrl: './album-list.html',
})
export class AlbumList extends BaseListPagination {

    public data: Album[];

    public listUrl = '/pages/albums';
    public title = 'Albums';

    constructor (
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        public service: AlbumService) {
        super();
    }

}