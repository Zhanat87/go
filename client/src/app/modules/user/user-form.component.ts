import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs/Rx';

import { BaseForm } from '../../common/base/baseForm';

import { User } from './user';
import { UserService } from './user.service';

import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';

import {CropperSettings} from 'ng2-img-cropper';
import {Environment} from "../../common/environment";

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

    cropperData: any;
    cropperSettings: CropperSettings;

    constructor(
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        protected localStorageService: LocalStorageService,
        public service: UserService) {
        super();

        this.initCropper();
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

                this.setModelData();
                this.setBreadCrumbs();
            },
            error => this.errorMessage = <any>error);
    }

    getBreadCrumbTitle(): string {
        return this.editMode ? (this.currentUser.id == this.model.id ?
            'Profile' : 'User: ' + this.model.username) : 'Create new user';
    }

    setModelData(): void {
        if (this.model.phones) {
            let phones = JSON.parse(this.model.phones);
            this.model.phone = phones.phone;
            this.model.mobile_phone = phones.mobile_phone;
        }
    }

    onAfterSave(model?: any): void {
        super.onAfterSave(model);
        this.updateIfNeedCurrentAdmin();
    }

    updateIfNeedCurrentAdmin(): void {
        if (this.editMode && this.currentUser.id == this.model.id) {
            this.localStorageService.set('currentUser', JSON.stringify(this.model));
            this._state.notifyChanged('currentUserUpdated');
        }
    }

    initCropper(): void {
        this.cropperSettings = new CropperSettings();
        this.cropperSettings.width = 100;
        this.cropperSettings.height = 100;
        this.cropperSettings.croppedWidth = 100;
        this.cropperSettings.croppedHeight = 100;
        this.cropperSettings.canvasWidth = 400;
        this.cropperSettings.canvasHeight = 300;

        this.cropperData = {};
    }

    fillModel(): void {
        if (this.cropperData.image) {
            this.model.avatar = this.cropperData.image;
        }
        this.model.phones = JSON.stringify({"phone": this.model.phone, "mobile_phone": this.model.mobile_phone});
        delete this.model.phone;
        delete this.model.mobile_phone;
    }

    deleteAvatar(event): void {
        this.model.avatar = null;
        this.model.avatar_string = null;
        this.service.update(this.model, this.model.id)
            .subscribe(
                data => {
                    if (data.id) {
                        this.updateIfNeedCurrentAdmin();

                        let target = event.currentTarget || event.target || event.srcElement;
                        jQuery(target).parent().remove();
                    } else {
                        console.log('error delete avatar', data);
                    }
                },
                error => this.errorMessage = <any>error);
    }

    get avatar(): string {
        return this.model.avatar_string ? (this.model.avatar_string.substr(0, 4) == 'http' ? this.model.avatar_string :
            Environment.API_ENDPOINT + 'static/users/avatars/' + this.model.avatar_string) : '';
    }

}