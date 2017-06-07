import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import {NewsReplicationService} from "./newsReplication.service";
import {BaseListPagination} from "../../common/base/baseListPagination";
import {NewsReplication} from "./newsReplication";
import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';

@Component({
    moduleId: 'newsReplication',
    selector: 'newsReplication',
    styleUrls: ['./newsReplication-list.scss'],
    templateUrl: './newsReplication-list.html',
})
export class NewsReplicationList extends BaseListPagination {

    public data: NewsReplication[];

    public listUrl = '/replication-news';
    public title = 'Replication news';

    constructor (
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        protected localStorageService: LocalStorageService,
        public service: NewsReplicationService) {
        super();
    }

}