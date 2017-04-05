import { Component }    from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs/Rx';

import { BaseForm } from '../../common/base/baseForm';

import { User }        from './user';
import { UserService } from './user.service';

import {GlobalState} from "../../global.state";

@Component({
    moduleId: 'user',
    selector: 'user-form',
    templateUrl: './form.html',
    styleUrls: ['./../../common/styles/form.scss'],
})

export class UserFormComponent extends BaseForm {

    public listUrl = '/pages/users';
    public title = 'Users';

    public model = new User();

    public dialCode: number;

    constructor(
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        public service: UserService) {
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

                this.model = data[0] as User;

                this.setBreadCrumbs();
            },
            error => this.errorMessage = <any>error);
    }

    onChangeCountry(event: any) {

    }

    reInitPlugins() {
    }

    getBreadCrumbTitle(): string {
        return this.editMode ? 'User: ' + this.model.name : 'Create new user';
    }

}