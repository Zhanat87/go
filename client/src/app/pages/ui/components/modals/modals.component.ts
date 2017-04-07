import {Component, ViewChild} from '@angular/core';
import {ModalDirective} from 'ngx-bootstrap';

@Component({
    selector: 'modals',
    styleUrls: ['./modals.scss'],
    templateUrl: './modals.html'
})
export class Modals {
    @ViewChild('childModal') childModal: ModalDirective;

    showChildModal(): void {
        this.childModal.show();
    }

    hideChildModal(): void {
        this.childModal.hide();
    }
}
