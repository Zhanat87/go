import { Component, EventEmitter, Input, Output } from '@angular/core';
import {ModalDirective} from 'ngx-bootstrap/modal';
import { ViewChild } from '@angular/core';
import {LocalStorageService} from "angular-2-local-storage";

@Component({
    selector: 'delete-modal',
    styleUrls: ['./delete.scss'],
    templateUrl: './delete.html',
})
export class CommonDeleteModalComponent {

    @ViewChild('deleteModal') deleteModal: ModalDirective;

    @Input() entity: string;
    @Output() deleteEntity: EventEmitter<string> = new EventEmitter<string>();

    constructor(protected localStorageService: LocalStorageService) {
    }

    showDeleteModal(id: number): void {
        this.localStorageService.set('deleteRowId', id);
        this.deleteModal.show();
    }

    hideDeleteModal(): void {
        this.localStorageService.remove('deleteRowId');
        this.deleteModal.hide();
    }

    delete() {
        this.deleteEntity.emit();
    }

}