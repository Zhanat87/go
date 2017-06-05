import {NgModule} from '@angular/core';
import {SharedModule} from "../../common/modules/shared.module";

import {NewsList} from './news-list.component';
import {NewsFormComponent} from './news-form.component';
import {NewsRoutingModule} from './news-routing.module';
import {NewsService} from "./news.service";

@NgModule({
    imports: [
        SharedModule,
        NewsRoutingModule,
    ],
    declarations: [
        NewsList,
        NewsFormComponent,
    ],
    providers: [
        NewsService,
    ],
})
export class NewsModule {}
