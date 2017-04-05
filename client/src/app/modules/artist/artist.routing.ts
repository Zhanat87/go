import {Routes, RouterModule}  from '@angular/router';
import {ModuleWithProviders} from '@angular/core';

import {ArtistList} from './artist-list.component';
import {ArtistFormComponent} from './artist-form.component';

// noinspection TypeScriptValidateTypes
export const routes: Routes = [
    { path: '', component: ArtistList },
    { path: 'create', component: ArtistFormComponent },
    { path: ':id', component: ArtistFormComponent }
];

export const routing: ModuleWithProviders = RouterModule.forChild(routes);
