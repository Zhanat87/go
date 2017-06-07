import {RouterModule} from '@angular/router';
import {NgModule} from '@angular/core';

import {NewsReplicationList} from './newsReplication-list.component';
import {NewsReplicationFormComponent} from './newsReplication-form.component';

@NgModule({
    imports: [RouterModule.forChild([
        { path: '',       component: NewsReplicationList },
        { path: 'create', component: NewsReplicationFormComponent },
        { path: ':id',    component: NewsReplicationFormComponent }
    ])],
    exports: [ RouterModule ]
})
export class NewsReplicationRoutingModule {}
