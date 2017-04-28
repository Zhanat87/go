import {RouterModule} from '@angular/router';
import {NgModule} from '@angular/core';

import {AlbumList} from './album-list.component';
import {AlbumFormComponent} from './album-form.component';

@NgModule({
    imports: [RouterModule.forChild([
        { path: '',       component: AlbumList },
        { path: 'create', component: AlbumFormComponent },
        { path: ':id',    component: AlbumFormComponent }
    ])],
    exports: [ RouterModule ]
})
export class AlbumRoutingModule {}
