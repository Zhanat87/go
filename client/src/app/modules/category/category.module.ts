import {NgModule} from '@angular/core';
import {SharedModule} from "../../common/modules/shared.module";

import {CategoryList} from './category-list.component';
import {CategoryFormComponent} from './category-form.component';
import {CategoryRoutingModule} from './category-routing.module';
import {CategoryService} from "./category.service";

@NgModule({
    imports: [
        SharedModule,
        CategoryRoutingModule,
    ],
    declarations: [
        CategoryList,
        CategoryFormComponent,
    ],
    providers: [
        CategoryService,
    ],
})
export class CategoryModule {}
