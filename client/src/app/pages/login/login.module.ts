import { NgModule }      from '@angular/core';
import { CommonModule }  from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgaModule } from '../../theme/nga.module';

import { Login } from './login.component';
import { routing }       from './login.routing';
import {LoginService} from "./login.service";
// import { provideAuth } from 'angular2-jwt';

@NgModule({
  imports: [
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
    NgaModule,
    routing,
  ],
  declarations: [
    Login,
  ],
  providers: [
    LoginService,
    // provideAuth(),
  ],
})
// @todo common AuthModule: login, logout, jwt refresh
export class LoginModule {}
