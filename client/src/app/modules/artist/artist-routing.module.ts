import {RouterModule} from '@angular/router';
import {NgModule} from '@angular/core';

import {ArtistList} from './artist-list.component';
import {ArtistFormComponent} from './artist-form.component';

@NgModule({
    imports: [RouterModule.forChild([
        { path: '',       component: ArtistList },
        { path: 'create', component: ArtistFormComponent },
        { path: ':id',    component: ArtistFormComponent }
    ])],
    exports: [ RouterModule ]
})
export class ArtistRoutingModule {}
