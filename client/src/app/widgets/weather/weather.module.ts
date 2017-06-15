import {NgModule} from '@angular/core';
import {CommonModule} from "@angular/common";
import { ModalModule } from 'ngx-bootstrap';

import {WeatherComponent} from './weather.component';
import {WeatherService} from "./weather.service";

@NgModule({
    imports: [
        CommonModule,
        ModalModule,
    ],
    declarations: [
        WeatherComponent,
    ],
    providers: [
        WeatherService,
    ],
    exports: [
        WeatherComponent,
    ],
})
export class WeatherModule {}
