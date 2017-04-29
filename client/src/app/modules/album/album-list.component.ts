import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import {AlbumService} from "./album.service";
import {BaseListPagination} from "../../common/base/baseListPagination";
import {Album} from "./album";
import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';

@Component({
    moduleId: 'album',
    selector: 'album',
    styleUrls: ['./album-list.scss'],
    templateUrl: './album-list.html',
})
export class AlbumList extends BaseListPagination {

    public data: Album[];

    public listUrl = '/albums';
    public title = 'Albums';

    constructor (
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        protected localStorageService: LocalStorageService,
        public service: AlbumService) {
        super();
        console.log('test');
    }

}