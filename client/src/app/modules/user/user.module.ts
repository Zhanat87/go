import {NgModule} from '@angular/core';
import {SharedModule} from "../../common/modules/shared.module";

import {UserList} from './user-list.component';
import {UserFormComponent} from './user-form.component';
import {UserRoutingModule} from './user-routing.module';
import {UserService} from "./user.service";

import {ImageCropperModule} from 'ng2-img-cropper';

@NgModule({
    imports: [
        SharedModule,
        UserRoutingModule,
        ImageCropperModule,
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
