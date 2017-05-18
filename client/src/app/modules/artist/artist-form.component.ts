import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs/Rx';

import { BaseForm } from '../../common/base/baseForm';

import { Artist } from './artist';
import { ArtistService } from './artist.service';

import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';
import {Environment} from "../../common/environment";

@Component({
    moduleId: 'artist',
    selector: 'artist-form',
    templateUrl: './form.html',
    styleUrls: ['./../../common/styles/form.scss'],
})

export class ArtistFormComponent extends BaseForm {

    public listUrl = '/artists';
    public title = 'Artists';

    public model = new Artist();

    constructor(
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        protected localStorageService: LocalStorageService,
        public service: ArtistService) {
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

                this.model = data[0] as Artist;

                this.setBreadCrumbs();
            },
            error => this.errorMessage = <any>error);
    }

    getBreadCrumbTitle(): string {
        return this.editMode ? 'Artist: ' + this.model.name : 'Create new artist';
    }

    deleteImage(event): void {
        this.model.image_base_64 = 'remove';
        this.service.update(this.model, this.model.id)
            .subscribe(
                data => {
                    if (data.id) {
                        let target = event.currentTarget || event.target || event.srcElement;
                        jQuery(target).parent().remove();
                        this.model.image = null;
                    } else {
                        console.log('error delete image', data);
                    }
                },
                error => this.errorMessage = <any>error);
    }

    onChangeImage(event: EventTarget) {
        let eventObj: MSInputMethodContext = <MSInputMethodContext>event;
        let target: HTMLInputElement = <HTMLInputElement>eventObj.target;
        let files: FileList = target.files;

        let reader = new FileReader();
        reader.readAsDataURL(files[0]);
        reader.onload = function () {
            jQuery(document.getElementById('image')).text(reader.result);
        };
        reader.onerror = function (error) {
            jQuery(document.getElementById('image')).text('');
        };
    }

    fillModel(): void {
        let imageBase64 = jQuery(document.getElementById('image')).text();
        this.model.image_base_64 = imageBase64 ? imageBase64 : null;
    }

    get image(): string {
        return this.model.image ? Environment.API_ENDPOINT + 'static/artists/images/' + this.model.image : '';
    }

}