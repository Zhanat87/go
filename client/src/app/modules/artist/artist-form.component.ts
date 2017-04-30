import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs/Rx';

import { BaseForm } from '../../common/base/baseForm';

import { Artist } from './artist';
import { ArtistService } from './artist.service';

import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';

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

    public dialCode: number;

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

}