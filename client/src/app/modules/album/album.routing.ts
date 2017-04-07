import {Routes, RouterModule}  from '@angular/router';
import {ModuleWithProviders} from '@angular/core';

import {AlbumList} from './album-list.component';
import {AlbumFormComponent} from './album-form.component';

// noinspection TypeScriptValidateTypes
export const routes: Routes = [
    { path: '', component: AlbumList },
    { path: 'create', component: AlbumFormComponent },
    { path: ':id', component: AlbumFormComponent }
];

export const routing: ModuleWithProviders = RouterModule.forChild(routes);
