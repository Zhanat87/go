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

    public dialCode: number;

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
                console.log(this.artists);

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

}