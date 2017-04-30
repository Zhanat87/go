import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {NgaModule} from '../../theme/nga.module';
import {Ng2PaginationModule} from 'ng2-pagination';
import { ModalModule } from 'ngx-bootstrap';
import { FormsModule } from '@angular/forms';
import {RouterModule} from "@angular/router";

// note: http or authHttp module include here

import {CommonPaginationFooterComponent} from "../components/pagination/footer/footer";
import {CommonPaginationHeaderComponent} from "../components/pagination/header/header";

// import { IterablePipe } from "../pipes/iterable";

@NgModule({
    imports: [
        CommonModule,
        NgaModule,
        Ng2PaginationModule,
        ModalModule.forRoot(),
        FormsModule,
        RouterModule,
    ],
    declarations: [
        CommonPaginationFooterComponent,
        CommonPaginationHeaderComponent,
        // IterablePipe,
    ],
    exports: [
        CommonModule,
        NgaModule,
        Ng2PaginationModule,
        ModalModule,
        FormsModule,
        RouterModule,

        CommonPaginationFooterComponent,
        CommonPaginationHeaderComponent,
        // IterablePipe,
    ],
})
export class SharedModule {}
