import {Component, ViewChild, OnInit} from '@angular/core';
import {WeatherService} from "./weather.service";
import {Weather} from "./weather";
import {ModalDirective} from 'ngx-bootstrap/modal';
import {setCurrentPosition} from "../../common/utils";

@Component({
    selector: 'weather',
    templateUrl: './weather.html',
    styleUrls: ['./weather.scss'],
})
export class WeatherComponent implements OnInit {

    @ViewChild('weatherModal') weatherModal: ModalDirective;

    public weather: Weather;

    constructor(private service: WeatherService) {
    }

    ngOnInit(): void {
        let $weatherModal = jQuery(document.getElementById('weatherModal')).remove();
        jQuery(document.getElementsByClassName('al-content')[0]).append($weatherModal);
        setCurrentPosition();
    }

    showWeatherInfo(): void {
        this.service.getWeatherInfo()
          .subscribe(
            data => {
              this.weather = data as Weather;
              this.showWeatherModal();
            },
            error => {
              console.log('showWeatherInfo error', error);
            },
          );
    }

    showWeatherModal(): void {
        this.weatherModal.show();
    }

    hideWeatherModal(): void {
        this.weatherModal.hide();
    }

}
