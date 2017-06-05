import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs/Rx';

import { BaseForm } from '../../common/base/baseForm';

import { Category } from './category';
import { CategoryService } from './category.service';

import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';

@Component({
    moduleId: 'category',
    selector: 'category-form',
    templateUrl: './form.html',
    styleUrls: ['./../../common/styles/form.scss'],
})

export class CategoryFormComponent extends BaseForm {

    public listUrl = '/categories';
    public title = 'Categories';

    public model = new Category();

    constructor(
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        protected localStorageService: LocalStorageService,
        public service: CategoryService) {
        super();
    }

    initCreateForm() {
        this.active = true;

        this.setBreadCrumbs();
    }

    initEditForm(id) {
        this.editMode = true;

        Observable.forkJoin(
            this.service.find(id)
        ).subscribe(
            data => {
                this.active = true;

                this.model = data[0] as Category;

                this.setBreadCrumbs();
            },
            error => this.errorMessage = <any>error);
    }

    getBreadCrumbTitle(): string {
        return this.editMode ? 'Category: ' + this.model.title : 'Create new category';
    }

}