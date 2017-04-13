import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';
import { Observable } from 'rxjs/Rx';

import { CommonService } from '../../common/services/service';
import { LogoutResponse } from './logout.response';

@Injectable()
export class LogoutService extends CommonService {
    public url = 'auth/invalidate';
    
    constructor (public http: AuthHttp) {
        super();
    }
    
    map(data): Observable<LogoutResponse> {
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
