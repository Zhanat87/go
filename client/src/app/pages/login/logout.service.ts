import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';
import { Observable } from 'rxjs/Rx';

import { CommonService } from '../../common/services/service';
import {SuccessResponse} from "../../common/entities/successResponse";

@Injectable()
export class LogoutService extends CommonService {
    public url = 'v1/auth/sign-out';
    
    constructor (public http: AuthHttp) {
        super();
    }
    
    map(data): Observable<SuccessResponse> {
        return data; 
    }

    public signOut() {
        return this.map(
            this.http.delete(this.getUrl())
                .map(this.extractAllData)
                .catch(this.handleError)
        );
    }
}
