import {Component} from '@angular/core';
import {FormGroup, AbstractControl, FormBuilder, Validators} from '@angular/forms';

import 'style-loader!./login.scss';
import {Router} from "@angular/router";
import {LoginService} from "./login.service";
import {LoginResponse} from "./login.response";

@Component({
    selector: 'login',
    templateUrl: './login.html',
})
export class Login {

    public form: FormGroup;
    public email: AbstractControl;
    public password: AbstractControl;

    public submitted: boolean = false;
    public error: boolean = false;
    public errorMessage: string;
    public active: boolean = true;

    public response: LoginResponse;

    constructor(fb: FormBuilder,
                private router: Router,
                private service: LoginService) {
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

                        if (this.response.data) {
                            localStorage.setItem('id_token', this.response.data.token);
                            localStorage.setItem('currentUser', JSON.stringify(this.response.data.user));
                            this.redirectToReferrer();
                        } else  {
                            this.error = true;
                            this.errorMessage = this.response.message;
                            this.email.valid = false;
                            this.password.valid = false;
                            this.email.touched = true;
                            this.password.touched = true;
                        }
                    },
                    error => {
                        this.errorMessage = <any>error
                    },
                    () => {
                        this.active = true;
                    },
                );
        }
    }

    redirectToReferrer(): void {
        this.router.navigate([localStorage.getItem('referrer') ? localStorage.getItem('referrer') : '/']);
    }

}
