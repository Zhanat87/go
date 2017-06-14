import {NgModule} from '@angular/core';
import {CommonModule} from "@angular/common";
import { ModalModule } from 'ngx-bootstrap';

import {CurrencyComponent} from './currency.component';
import {CurrencyService} from "./currency.service";

@NgModule({
    imports: [
        CommonModule,
        ModalModule,
    ],
    declarations: [
        CurrencyComponent,
    ],
    providers: [
        CurrencyService,
    ],
    exports: [
        CurrencyComponent,
    ],
})
export class CurrencyModule {}
