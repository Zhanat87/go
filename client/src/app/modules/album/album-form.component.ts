import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs/Rx';

import { BaseForm } from '../../common/base/baseForm';

import { Album } from './album';
import { AlbumService } from './album.service';

import { Artist } from '../artist/artist';
import { ArtistService } from '../artist/artist.service';

import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';
import {Environment} from "../../common/environment";

@Component({
    moduleId: 'album',
    selector: 'album-form',
    templateUrl: './form.html',
    styleUrls: ['./../../common/styles/form.scss'],
    providers: [
        ArtistService,
    ],
})

export class AlbumFormComponent extends BaseForm {

    public listUrl = '/albums';
    public title = 'Albums';

    public model = new Album();

    public artists: Artist[];

    constructor(
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        protected localStorageService: LocalStorageService,
        private artistService: ArtistService,
        public service: AlbumService) {
        super();
    }

    initCreateForm() {
        Observable.forkJoin(
            this.artistService.allWithoutLimit()
        ).subscribe(
            data => {
                this.active = true;

                this.artists = data[0] as Artist[];

                this.setBreadCrumbs();
            },
            error => this.errorMessage = <any>error);
    }

    initEditForm(id) {
        this.editMode = true;

        Observable.forkJoin(
            this.service.find(id),
            this.artistService.allWithoutLimit()
        ).subscribe(
            data => {
                this.active = true;

                this.model = data[0] as Album;
                this.artists = data[1] as Artist[];

                this.setBreadCrumbs();
            },
            error => this.errorMessage = <any>error);
    }

    getBreadCrumbTitle(): string {
        return this.editMode ? 'Album: ' + this.model.title : 'Create new album';
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
        return this.model.image ? Environment.AWS_S3_BUCKET_URL + 'static/albums/images/' +
            this.getThumbPath(this.model.image, 100) : '';
    }

    getThumbPath (image: string, size: number): string {
        return image.substr(0, image.indexOf('.')) + `/${size}_` + image;
    }

}