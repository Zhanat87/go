import {Injectable} from '@angular/core';
import {Http} from '@angular/http';
import {Observable} from 'rxjs/Rx';

import {CommonService} from '../../common/services/service';
import {SuccessResponse} from "../../common/entities/successResponse";

@Injectable()
export class PasswordResetRequestService extends CommonService {

    public url = 'v1/auth/password-reset-request';

    constructor(public http: Http) {
        super();
    }

    map(data): Observable<SuccessResponse> {
        return data;
    }

    public passwordResetRequest(attributes: any) {
        return this.map(
            this.http.post(this.getUrl(), attributes)
                .map(this.extractAllData)
                .catch(this.handleServerErrors)
        );
    }

}