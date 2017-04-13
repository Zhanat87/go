import {NgModule}      from '@angular/core';
import {CommonModule}  from '@angular/common';

import {routing}       from './pages.routing';
import {NgaModule} from '../theme/nga.module';

import {Pages} from './pages.component';
import {AuthGuard} from "../common/services/auth.guard";

@NgModule({
    imports: [
        CommonModule,
        NgaModule,
        routing,
    ],
    declarations: [
        Pages,
    ],
    providers: [
        AuthGuard,
    ],
})
export class PagesModule {}
