import {NgModule} from '@angular/core';
import { Title } from '@angular/platform-browser';

import {AppState} from "../../app.service";
import {GlobalState} from "../../global.state";
import {LoginService} from "../../pages/login/login.service";

import { LocalStorageModule } from 'angular-2-local-storage';

/*
 * Platform and Environment providers/directives/pipes
 */
import {ENV_PROVIDERS} from "../../environment";

// common services must be here
@NgModule({
    providers:    [
        ENV_PROVIDERS,
        AppState,
        GlobalState,
        LoginService,
        Title,
    ],
    imports: [
        LocalStorageModule.withConfig({
            prefix: '', // my-app
            storageType: 'localStorage'
        }),
    ],
})
export class CoreModule {}