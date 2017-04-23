import { Component }    from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs/Rx';

import { BaseForm } from '../../common/base/baseForm';

import { User }        from './user';
import { UserService } from './user.service';

import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';

@Component({
    moduleId: 'user',
    selector: 'user-form',
    templateUrl: './form.html',
    styleUrls: ['./../../common/styles/form.scss'],
})

export class UserFormComponent extends BaseForm {

    public listUrl = '/users';
    public title = 'Users';

    public model = new User();

    private currentUser: User;

    constructor(
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        protected localStorageService: LocalStorageService,
        public service: UserService) {
        super();
    }

    initCreateForm() {

        this.active = true;

        this.setBreadCrumbs();

    }

    initEditForm(id) {
        this.editMode = true;

        this.currentUser = this.getCurrentUser();

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

    getBreadCrumbTitle(): string {
        return this.editMode ? (this.currentUser.id == this.model.id ? 'Profile' : 'User: ' + this.model.username) : 'Create new user';
    }

    onAfterSave(): void {
        if (this.editMode && this.currentUser.id == this.model.id) {
            this.localStorageService.set('currentUser', JSON.stringify({id: this.model.id, username: this.model.username, email: this.model.email}));
            this._state.notifyChanged('currentUserUpdated');
        }
        super.onAfterSave();
    }

}