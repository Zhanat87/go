import {RouterModule} from '@angular/router';
import {NgModule} from '@angular/core';

import {UserList} from './user-list.component';
import {UserFormComponent} from './user-form.component';

@NgModule({
    imports: [RouterModule.forChild([
        { path: '',       component: UserList },
        { path: 'create', component: UserFormComponent },
        { path: ':id',    component: UserFormComponent }
    ])],
    exports: [ RouterModule ]
})
export class UserRoutingModule {}
