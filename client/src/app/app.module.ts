import {NgModule, ApplicationRef} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {HttpModule} from '@angular/http';
import {RouterModule} from '@angular/router';
import {removeNgStyles, createNewHosts, createInputTransfer} from '@angularclass/hmr';

// Routing
import { AppRoutingModule } from './app-routing.module';

// App is our top level component
import {App} from './app.component';
import {AppState, InternalStateType} from './app.service';
import {NgaModule} from './theme/nga.module';

// Importing modules
import { AlbumModule } from "./modules/album/album.module";
import { UserModule } from "./modules/user/user.module";
import { ArtistModule } from "./modules/artist/artist.module";
import { AuthModule } from './modules/auth/auth.module';
import { CategoryModule } from "./modules/category/category.module";
import { NewsModule } from "./modules/news/news.module";

// component's without module
import { PageComponents } from "./pages/pages.components";
import { MainLayoutComponent } from "./common/components/layouts/main/main.component";
import { BlankLayoutComponent } from "./common/components/layouts/blank/blank.component";

import {CoreModule} from "./common/modules/core.module";

export type StoreType = {
    state: InternalStateType,
    restoreInputValues: () => void,
    disposeOldHosts: () => void
};

/**
 * `AppModule` is the main entry point into Angular2's bootstraping process
 */
@NgModule({
    bootstrap: [App],
    declarations: [
        App,
        ...PageComponents,
        MainLayoutComponent,
        BlankLayoutComponent,
    ],
    imports: [ // import Angular's modules
        BrowserModule,
        HttpModule,
        RouterModule,
        FormsModule,
        ReactiveFormsModule,
        NgaModule.forRoot(),

        CoreModule,

        // note: must be at head of all modules with routing modules
        AppRoutingModule,

        // app modules
        AuthModule,
        AlbumModule,
        UserModule,
        ArtistModule,
        CategoryModule,
        // NewsModule,
    ],
})

export class AppModule {

    constructor(public appRef: ApplicationRef, public appState: AppState) {
    }

    hmrOnInit(store: StoreType) {
        if (!store || !store.state) return;
        console.log('HMR store', JSON.stringify(store, null, 2));
        // set state
        this.appState._state = store.state;
        // set input values
        if ('restoreInputValues' in store) {
            let restoreInputValues = store.restoreInputValues;
            setTimeout(restoreInputValues);
        }
        this.appRef.tick();
        delete store.state;
        delete store.restoreInputValues;
    }

    hmrOnDestroy(store: StoreType) {
        const cmpLocation = this.appRef.components.map(cmp => cmp.location.nativeElement);
        // save state
        const state = this.appState._state;
        store.state = state;
        // recreate root elements
        store.disposeOldHosts = createNewHosts(cmpLocation);
        // save input values
        store.restoreInputValues = createInputTransfer();
        // remove styles
        removeNgStyles();
    }

    hmrAfterDestroy(store: StoreType) {
        // display new elements
        store.disposeOldHosts();
        delete store.disposeOldHosts;
    }
}
