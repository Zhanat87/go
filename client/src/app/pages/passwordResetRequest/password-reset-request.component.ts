import {Component} from '@angular/core';
import {FormGroup, AbstractControl, FormBuilder, Validators} from '@angular/forms';
import {EmailValidator} from '../../theme/validators';
import {PasswordResetRequestService} from "./password-reset-request.service";
import {SuccessResponse} from "../../common/entities/successResponse";

@Component({
    selector: 'password-reset-request',
    templateUrl: './password-reset-request.html',
    styleUrls: ['./password-reset-request.scss'],
    providers: [
        PasswordResetRequestService,
    ],
})
export class PasswordResetRequestComponent {

    public form: FormGroup;
    public email: AbstractControl;

    public submitted: boolean = false;

    public errorMessage: string;
    private response: SuccessResponse;

    public message: string;

    constructor(fb: FormBuilder,
                private service: PasswordResetRequestService) {

        this.form = fb.group({
            'email': ['', Validators.compose([Validators.required, EmailValidator.validate])],
        });

        this.email = this.form.controls['email'];
    }

    public onSubmit(values: any): void {
        this.submitted = true;
        if (this.form.valid) {
            this.service.passwordResetRequest(values)
                .subscribe(
                    data => {
                        this.response = data as SuccessResponse;
                        if (this.response.status == 200) {
                            this.message = this.response.message;
                        }
                    },
                    error => {
                        this.handleErrors(error);
                    },
                );
        }
    }

    handleErrors(error: string): void {
        switch (error) {
            case 'the requested resource was not found.':
                this.errorMessage = 'user with this email not exist';
                this.email.valid = false;
                break;
        }
    }

}