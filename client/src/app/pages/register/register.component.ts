import {Component} from '@angular/core';
import {FormGroup, AbstractControl, FormBuilder, Validators} from '@angular/forms';
import {EmailValidator, EqualPasswordsValidator} from '../../theme/validators';
import {RegisterService} from "./register.service";
import {SuccessResponse} from "../../common/entities/successResponse";
import {Router} from "@angular/router";

@Component({
    selector: 'register',
    templateUrl: './register.html',
    styleUrls: ['./register.scss'],
    providers: [
        RegisterService,
    ],
})
export class RegisterComponent {

    public form: FormGroup;
    public name: AbstractControl;
    public email: AbstractControl;
    public password: AbstractControl;
    public repeatPassword: AbstractControl;
    public passwords: FormGroup;

    public submitted: boolean = false;

    public errorMessage: string;
    private response: SuccessResponse;

    constructor(fb: FormBuilder,
                private router: Router,
                private service: RegisterService) {

        this.form = fb.group({
            'name': ['', Validators.compose([Validators.required, Validators.minLength(4)])],
            'email': ['', Validators.compose([Validators.required, EmailValidator.validate])],
            'passwords': fb.group({
                'password': ['', Validators.compose([Validators.required, Validators.minLength(4)])],
                'repeatPassword': ['', Validators.compose([Validators.required, Validators.minLength(4)])]
            }, {validator: EqualPasswordsValidator.validate('password', 'repeatPassword')})
        });

        this.name = this.form.controls['name'];
        this.email = this.form.controls['email'];
        this.passwords = <FormGroup> this.form.controls['passwords'];
        this.password = this.passwords.controls['password'];
        this.repeatPassword = this.passwords.controls['repeatPassword'];
    }

    public onSubmit(values: any): void {
        this.submitted = true;
        if (this.form.valid) {
            values.password = values.passwords.password;
            delete values.passwords;
            this.service.signUp(values)
                .subscribe(
                    data => {
                        this.response = data as SuccessResponse;
                        if (this.response.status == 200) {
                            this.redirectToLogin();
                        }
                    },
                    error => {
                        this.handleErrors(error);
                    },
                );
        }
    }

    redirectToLogin(): void {
        this.router.navigate(['/login']);
    }

    handleErrors(error: string): void {
        switch (error) {
            case 'Internal server error: pq: duplicate key value violates unique constraint "constraint_unique_email_by_provider"':
                this.errorMessage = 'user with this email exist';
                this.email.valid = false;
                break;
        }
    }

}