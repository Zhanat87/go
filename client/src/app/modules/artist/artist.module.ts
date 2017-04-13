import {NgModule}      from '@angular/core';
import {SharedModule} from "../../common/modules/shared.module";

import {ArtistList} from './artist-list.component';
import {ArtistFormComponent} from './artist-form.component';
import {routing}       from './artist.routing';
import {ArtistService} from "./artist.service";

@NgModule({
    imports: [
        SharedModule,
        routing,
    ],
    declarations: [
        ArtistList,
        ArtistFormComponent,
    ],
    providers: [
        ArtistService,
    ],
})
export class ArtistModule {}
