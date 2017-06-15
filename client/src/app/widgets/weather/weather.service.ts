import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';
import { Observable } from 'rxjs/Rx';
import {Weather} from "./weather";
import { CommonService } from '../../common/services/service';

@Injectable()
export class WeatherService extends CommonService {
    public url = 'v1/weather-info';

    constructor (public http: AuthHttp) {
      super();
    }

    map(data): Observable<Weather> {
        return data;
    }

    public getWeatherInfo(): Observable<Weather> {
        let lat = localStorage.getItem('lat');
        let lon = localStorage.getItem('lon');
        return this.map(
            this.http.get(this.getUrl() + `/${lat}/${lon}`)
                .map(this.extractAllData)
                .catch(this.handleError)
        );
    }
}
