import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import {CategoryService} from "./category.service";
import {BaseListPagination} from "../../common/base/baseListPagination";
import {Category} from "./category";
import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';

@Component({
    moduleId: 'category',
    selector: 'category',
    styleUrls: ['./category-list.scss'],
    templateUrl: './category-list.html',
})
export class CategoryList extends BaseListPagination {

    public data: Category[];

    public listUrl = '/categories';
    public title = 'Categories';

    constructor (
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        protected localStorageService: LocalStorageService,
        public service: CategoryService) {
        super();
    }

}