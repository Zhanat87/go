import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import {NewsService} from "./news.service";
import {BaseListPagination} from "../../common/base/baseListPagination";
import {News} from "./news";
import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';

@Component({
    moduleId: 'news',
    selector: 'news',
    styleUrls: ['./news-list.scss'],
    templateUrl: './news-list.html',
})
export class NewsList extends BaseListPagination {

    public data: News[];

    public listUrl = '/news';
    public title = 'News';

    constructor (
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        protected localStorageService: LocalStorageService,
        public service: NewsService) {
        super();
    }

}