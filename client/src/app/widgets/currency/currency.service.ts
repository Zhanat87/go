import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';
import { Observable } from 'rxjs/Rx';
import {Currency} from "./currency";
import { CommonService } from '../../common/services/service';

@Injectable()
export class CurrencyService extends CommonService {
    public url = 'v1/currencies';

    constructor (public http: AuthHttp) {
      super();
    }

    map(data): Observable<Currency[]> {
        return data;
    }

    public getExchangeRates() {
        return this.map(
            this.http.get(this.getUrl())
                .map(this.extractData)
                .catch(this.handleError)
        );
    }
}
