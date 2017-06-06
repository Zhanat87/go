import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import {NewsShardService} from "./newsShard.service";
import {BaseListPagination} from "../../common/base/baseListPagination";
import {NewsShard} from "./newsShard";
import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';

@Component({
    moduleId: 'newsShard',
    selector: 'newsShard',
    styleUrls: ['./newsShard-list.scss'],
    templateUrl: './newsShard-list.html',
})
export class NewsShardList extends BaseListPagination {

    public data: NewsShard[];

    public listUrl = '/shard-news';
    public title = 'Shard news';

    constructor (
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        protected localStorageService: LocalStorageService,
        public service: NewsShardService) {
        super();
    }

}