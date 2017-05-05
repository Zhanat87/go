import {NgModule} from '@angular/core';
import {SharedModule} from "../../common/modules/shared.module";

import {UserList} from './user-list.component';
import {UserFormComponent} from './user-form.component';
import {UserRoutingModule} from './user-routing.module';
import {UserService} from "./user.service";

import {ImageCropperComponent} from 'ng2-img-cropper';

@NgModule({
    imports: [
        SharedModule,
        UserRoutingModule,
    ],
    declarations: [
        UserList,
        UserFormComponent,
        ImageCropperComponent,
    ],
    providers: [
        UserService,
    ],
})
export class UserModule {}
