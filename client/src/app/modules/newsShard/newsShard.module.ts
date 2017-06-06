import {NgModule} from '@angular/core';
import {SharedModule} from "../../common/modules/shared.module";

import {NewsShardList} from './newsShard-list.component';
import {NewsShardFormComponent} from './newsShard-form.component';
import {NewsShardRoutingModule} from './newsShard-routing.module';
import {NewsShardService} from "./newsShard.service";

@NgModule({
    imports: [
        SharedModule,
        NewsShardRoutingModule,
    ],
    declarations: [
        NewsShardList,
        NewsShardFormComponent,
    ],
    providers: [
        NewsShardService,
    ],
})
export class NewsShardModule {}
