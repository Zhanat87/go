import { Component }    from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs/Rx';

import { BaseForm } from '../../common/base/baseForm';

import { Album }        from './album';
import { AlbumService } from './album.service';

import {GlobalState} from "../../global.state";

@Component({
    moduleId: 'album',
    selector: 'album-form',
    templateUrl: './form.html',
    styleUrls: ['./../../common/styles/form.scss'],
})

export class AlbumFormComponent extends BaseForm {

    public listUrl = '/pages/albums';
    public title = 'Albums';

    public model = new Album();

    public dialCode: number;

    constructor(
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        public service: AlbumService) {
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

                this.model = data[0] as Album;

                this.setBreadCrumbs();
            },
            error => this.errorMessage = <any>error);
    }

    onChangeCountry(event: any) {

    }

    getBreadCrumbTitle(): string {
        return this.editMode ? 'Album: ' + this.model.title : 'Create new album';
    }

}