import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {NgaModule} from '../../theme/nga.module';
import {NgxPaginationModule} from 'ngx-pagination';
import { ModalModule } from 'ngx-bootstrap';
import { FormsModule } from '@angular/forms';
import {RouterModule} from "@angular/router";

// note: http or authHttp module include here

import {CommonPaginationFooterComponent} from "../components/pagination/footer/footer";
import {CommonPaginationHeaderComponent} from "../components/pagination/header/header";

import { CKEditorComponent } from '../components/ckeditor/ckeditor.component';

@NgModule({
    imports: [
        CommonModule,
        NgaModule,
        NgxPaginationModule,
        ModalModule.forRoot(),
        FormsModule,
        RouterModule,
    ],
    declarations: [
        CommonPaginationFooterComponent,
        CommonPaginationHeaderComponent,
        CKEditorComponent,
    ],
    exports: [
        CommonModule,
        NgaModule,
        NgxPaginationModule,
        ModalModule,
        FormsModule,
        RouterModule,

        CommonPaginationFooterComponent,
        CommonPaginationHeaderComponent,
        CKEditorComponent,
    ],
})
export class SharedModule {}
