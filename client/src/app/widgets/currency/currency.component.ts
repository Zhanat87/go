import {Component, ViewChild, OnInit} from '@angular/core';
import {CurrencyService} from "./currency.service";
import {Currency} from "./currency";
import {ModalDirective} from 'ngx-bootstrap/modal';

@Component({
    selector: 'currency',
    templateUrl: './currency.html',
    styleUrls: ['./currency.scss'],
    providers: [
       CurrencyService,
    ],
})
export class CurrencyComponent implements OnInit {

    @ViewChild('currencyModal') currencyModal: ModalDirective;

    public currencies: Currency[];

    constructor(private service: CurrencyService) {
    }

    ngOnInit(): void {
        let $currencyModal = jQuery(document.getElementById('currencyModal')).remove();
        jQuery(document.getElementsByClassName('al-content')[0]).append($currencyModal);
    }

    showExchangeRates(): void {
        this.service.getExchangeRates()
          .subscribe(
            data => {
              this.currencies = data as Currency[];
              this.showCurrencyModal();
            },
            error => {
              console.log('showExchangeRates error', error);
            },
          );
    }

    showCurrencyModal(): void {
        this.currencyModal.show();
    }

    hideCurrencyModal(): void {
        this.currencyModal.hide();
    }

}
