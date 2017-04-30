import {Component} from '@angular/core';
import {FormGroup, AbstractControl, FormBuilder, Validators} from '@angular/forms';

import 'style-loader!./login.scss';
import {Router} from "@angular/router";
import {LoginService} from "./login.service";
import {LoginResponse} from "./login.response";
import { LocalStorageService } from 'angular-2-local-storage';

@Component({
    selector: 'login',
    templateUrl: './login.html',
})
export class LoginComponent {

    public form: FormGroup;
    public email: AbstractControl;
    public password: AbstractControl;

    public submitted: boolean = false;
    public error: boolean = false;
    public errorMessage: string;
    public active: boolean = true;

    public emailError: boolean = false;
    public passwordError: boolean = false;

    public response: LoginResponse;

    constructor(fb: FormBuilder,
                private router: Router,
                private service: LoginService,
                private localStorageService: LocalStorageService) {
        this.form = fb.group({
            'email': ['', Validators.compose([Validators.required, Validators.minLength(4)])],
            'password': ['', Validators.compose([Validators.required, Validators.minLength(4)])]
        });

        this.email = this.form.controls['email'];
        this.password = this.form.controls['password'];
    }

    public onSubmit(values: Object): void {
        this.submitted = true;
        this.active = true;
        if (this.form.valid) {
            this.error = false;
            this.errorMessage = '';

            this.service.signIn(values)
                .subscribe(
                    data => {
                        this.response = data as LoginResponse;

                        this.localStorageService.set('id_token', this.response.data.token);
                        this.localStorageService.set('currentUser', JSON.stringify(this.response.data.user));
                        this.redirectToReferrer();
                    },
                    error => {
                        this.serverErrors(error);
                    },
                    () => {
                        this.active = true;
                    },
                );
        }
    }

    serverErrors(error: any): void {
        this.error = true;
        this.errorMessage = error;

        if (error == 'Authentication failed: user with this username not found' ||
            error == 'Authentication failed: user with this email not found') {
            this.email.valid = false;
            this.email.touched = true;

            this.password.valid = true;
            this.password.touched = false;

            this.emailError = true;
            this.passwordError = false;
        } else if (error == 'Authentication failed: password not valid') {
            this.password.valid = false;
            this.password.touched = true;

            this.email.valid = true;
            this.email.touched = false;

            this.passwordError = true;
            this.emailError = false;
        }
    }

    redirectToReferrer(): void {
        this.router.navigate([this.localStorageService.get('referrer') ? this.localStorageService.get('referrer') : '/']);
    }

}
