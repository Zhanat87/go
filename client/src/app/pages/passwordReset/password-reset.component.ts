import {Component, OnDestroy, OnInit} from '@angular/core';
import {FormGroup, AbstractControl, FormBuilder, Validators} from '@angular/forms';
import {PasswordResetService} from "./password-reset.service";
import {SuccessResponse} from "../../common/entities/successResponse";
import {Subscription} from "rxjs/Subscription";
import {ActivatedRoute, Router} from "@angular/router";
import {EqualPasswordsValidator} from "../../theme/validators/equalPasswords.validator";

@Component({
    selector: 'password-reset',
    templateUrl: './password-reset.html',
    styleUrls: ['./password-reset.scss'],
    providers: [
        PasswordResetService,
    ],
})
export class PasswordResetComponent implements OnInit, OnDestroy {

    public form: FormGroup;
    public password: AbstractControl;
    public repeatPassword: AbstractControl;
    public passwords: FormGroup;

    public submitted: boolean = false;

    public errorMessage: string;
    private response: SuccessResponse;

    private sub: Subscription;
    private token: string;

    constructor(fb: FormBuilder,
                private route: ActivatedRoute,
                private router: Router,
                private service: PasswordResetService) {

        this.form = fb.group({
            'passwords': fb.group({
                'password': ['', Validators.compose([Validators.required, Validators.minLength(4)])],
                'repeatPassword': ['', Validators.compose([Validators.required, Validators.minLength(4)])]
            }, {validator: EqualPasswordsValidator.validate('password', 'repeatPassword')})
        });

        this.passwords = <FormGroup> this.form.controls['passwords'];
        this.password = this.passwords.controls['password'];
        this.repeatPassword = this.passwords.controls['repeatPassword'];
    }

    /**
     * @link https://angular-2-training-book.rangle.io/handout/routing/routeparams.html
     */
    ngOnInit(): void {
        this.sub = this.route.params.subscribe(params => {
            this.token = params['token'];
        });
    }

    ngOnDestroy() {
        this.sub.unsubscribe();
    }

    public onSubmit(values: any): void {
        this.submitted = true;
        if (this.form.valid) {
            values.password = values.passwords.password;
            delete values.passwords;
            this.service.passwordReset(this.token, values)
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
            case 'the requested resource was not found.':
                this.errorMessage = 'token not exist';
                this.password.valid = false;
                break;
        }
    }

}