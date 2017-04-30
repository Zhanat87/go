import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Rx';

import { CommonService } from '../../common/services/service';
import { LoginResponse } from './login.response';

@Injectable()
export class LoginService extends CommonService {
    public url = 'v1/auth';
    
    constructor (public http: Http) {
        super();
    }
    
    map(data): Observable<LoginResponse> {
        return data; 
    }

    public signIn(attributes: any) {
        return this.map(
            this.http.post(this.getUrl(), attributes)
                .map(this.extractAllData)
                .catch(this.handleUnauthorizedError)
        );
    }
}
