import {NgModule} from '@angular/core';
import {SharedModule} from "../../common/modules/shared.module";

import {UserList} from './user-list.component';
import {UserFormComponent} from './user-form.component';
import {UserRoutingModule} from './user-routing.module';
import {UserService} from "./user.service";

@NgModule({
    imports: [
        SharedModule,
        UserRoutingModule,
    ],
    declarations: [
        UserList,
        UserFormComponent,
    ],
    providers: [
        UserService,
    ],
})
export class UserModule {}
