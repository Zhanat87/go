import {NgModule} from '@angular/core';
import {SharedModule} from "../../common/modules/shared.module";

import {NewsReplicationList} from './newsReplication-list.component';
import {NewsReplicationFormComponent} from './newsReplication-form.component';
import {NewsReplicationRoutingModule} from './newsReplication-routing.module';
import {NewsReplicationService} from "./newsReplication.service";

@NgModule({
    imports: [
        SharedModule,
        NewsReplicationRoutingModule,
    ],
    declarations: [
        NewsReplicationList,
        NewsReplicationFormComponent,
    ],
    providers: [
        NewsReplicationService,
    ],
})
export class NewsReplicationModule {}
