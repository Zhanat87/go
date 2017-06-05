import {RouterModule} from '@angular/router';
import {NgModule} from '@angular/core';

import {CategoryList} from './category-list.component';
import {CategoryFormComponent} from './category-form.component';

@NgModule({
    imports: [RouterModule.forChild([
        { path: '',       component: CategoryList },
        { path: 'create', component: CategoryFormComponent },
        { path: ':id',    component: CategoryFormComponent }
    ])],
    exports: [ RouterModule ]
})
export class CategoryRoutingModule {}
