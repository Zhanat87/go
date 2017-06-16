import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';
import { Observable } from 'rxjs/Rx';
import {McdonaldsItem} from "./mcdonaldsItem";
import { CommonService } from '../../common/services/service';

@Injectable()
export class McdonaldsService extends CommonService {
    public url = 'v1/mcdonalds';

    constructor (public http: AuthHttp) {
      super();
    }

    map(data): Observable<McdonaldsItem[]> {
        return data;
    }

    public getMcdonaldsMenu() {
        return this.map(
            this.http.get(this.getUrl())
                .map(this.extractAllData)
                .catch(this.handleError)
        );
    }
}
