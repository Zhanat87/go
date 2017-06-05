import {RouterModule} from '@angular/router';
import {NgModule} from '@angular/core';

import {NewsList} from './news-list.component';
import {NewsFormComponent} from './news-form.component';

@NgModule({
    imports: [RouterModule.forChild([
        { path: '',       component: NewsList },
        { path: 'create', component: NewsFormComponent },
        { path: ':id',    component: NewsFormComponent }
    ])],
    exports: [ RouterModule ]
})
export class NewsRoutingModule {}
