import {RouterModule} from '@angular/router';
import {NgModule} from '@angular/core';

import {NewsShardList} from './newsShard-list.component';
import {NewsShardFormComponent} from './newsShard-form.component';

@NgModule({
    imports: [RouterModule.forChild([
        { path: '',       component: NewsShardList },
        { path: 'create', component: NewsShardFormComponent },
        { path: ':id',    component: NewsShardFormComponent }
    ])],
    exports: [ RouterModule ]
})
export class NewsShardRoutingModule {}
