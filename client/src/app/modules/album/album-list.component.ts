import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import {AlbumService} from "./album.service";
import {BaseListPagination} from "../../common/base/baseListPagination";
import {Album} from "./album";
import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';
import {Environment} from "../../common/environment";

@Component({
    moduleId: 'album',
    selector: 'album',
    styleUrls: ['./album-list.scss'],
    templateUrl: './album-list.html',
})
export class AlbumList extends BaseListPagination {

    public data: Album[];

    public listUrl = '/albums';
    public title = 'Albums';

    constructor (
        public router: Router,
        public route: ActivatedRoute,
        protected _state: GlobalState,
        protected localStorageService: LocalStorageService,
        public service: AlbumService) {
        super();
    }

    getImage(album: Album): string {
        return album.image ? Environment.AWS_S3_BUCKET_URL + 'static/albums/images/' + this.getThumbPath(album.image, 100) : '';
    }

    getThumbPath (image: string, size: number): string {
        return image.substr(0, image.indexOf('.')) + `/${size}_` + image;
    }

}