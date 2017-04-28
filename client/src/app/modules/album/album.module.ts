import {NgModule}      from '@angular/core';
import {SharedModule} from "../../common/modules/shared.module";

import {AlbumList} from './album-list.component';
import {AlbumFormComponent} from './album-form.component';
import {AlbumRoutingModule}       from './album-routing.module';
import {AlbumService} from "./album.service";

@NgModule({
    imports: [
        SharedModule,
        AlbumRoutingModule,
    ],
    declarations: [
        AlbumList,
        AlbumFormComponent,
    ],
    providers: [
        AlbumService,
    ],
})
export class AlbumModule {}
