import {Routes, RouterModule}  from '@angular/router';
import {ModuleWithProviders} from '@angular/core';

import {UserList} from './user-list.component';
import {UserFormComponent} from './user-form.component';

// noinspection TypeScriptValidateTypes
export const routes: Routes = [
    { path: '', component: UserList },
    { path: 'create', component: UserFormComponent },
    { path: ':id', component: UserFormComponent }
];

export const routing: ModuleWithProviders = RouterModule.forChild(routes);
