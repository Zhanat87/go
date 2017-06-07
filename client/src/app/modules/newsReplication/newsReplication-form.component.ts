import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs/Rx';

import { BaseForm } from '../../common/base/baseForm';

import { NewsReplication } from './newsReplication';
import { NewsReplicationService } from './newsReplication.service';

import { Category } from '../category/category';
import { CategoryService } from '../category/category.service';

import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';

@Component({
    moduleId: 'newsReplication',
    selector: 'newsReplication-form',
    templateUrl: './form.html',
    styleUrls: ['./../../common/styles/form.scss'],
    providers: [
        CategoryService,
    ],
})

export class NewsReplicationFormComponent extends BaseForm {

    public listUrl = '/replication-news';
    public title = 'Replication news';

    public model = new NewsReplication();

    public categories: Category[];

    constructor(
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        protected localStorageService: LocalStorageService,
        private categoryService: CategoryService,
        public service: NewsReplicationService) {
        super();
    }

    initCreateForm() {
        Observable.forkJoin(
            this.categoryService.allWithoutLimit()
        ).subscribe(
            data => {
                this.active = true;

                this.categories = data[0] as Category[];

                this.setBreadCrumbs();
            },
            error => this.errorMessage = <any>error);
    }

    initEditForm(id) {
        this.editMode = true;

        Observable.forkJoin(
            this.service.find(id),
            this.categoryService.allWithoutLimit()
        ).subscribe(
            data => {
                this.active = true;

                this.model = data[0] as NewsReplication;
                this.categories = data[1] as Category[];

                this.setBreadCrumbs();
            },
            error => this.errorMessage = <any>error);
    }

    getBreadCrumbTitle(): string {
        return this.editMode ? 'Replication news: ' + this.model.title : 'Create new replication news';
    }

}