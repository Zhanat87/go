import {CommonListPagination} from "../list/pagination";
import {ModalDirective} from 'ng2-bootstrap';
import { ViewChild } from '@angular/core';

export abstract class BaseListPagination extends CommonListPagination {

    @ViewChild('deleteModal') deleteModal: ModalDirective;
    @ViewChild('unPublishModal') unPublishModal: ModalDirective;
    @ViewChild('publishModal') publishModal: ModalDirective;

    ngOnDestroy() {
        super.ngOnDestroy();

        localStorage.removeItem('deleteRowId');
        localStorage.removeItem('unPublishOfferId');
        localStorage.removeItem('publishOfferId');
    }

    showDeleteModal(id: number): void {
        localStorage.setItem('deleteRowId', id.toString());
        this.deleteModal.show();
    }

    hideDeleteModal(): void {
        localStorage.removeItem('deleteRowId');
        this.deleteModal.hide();
    }

    showUnPublishModal(id: number): void {
        localStorage.setItem('unPublishOfferId', id.toString());
        this.unPublishModal.show();
    }

    hideUnPublishModal(): void {
        localStorage.removeItem('unPublishOfferId');
        this.unPublishModal.hide();
    }

    unPublish(): void {
        this.service.unPublish(parseInt(localStorage.getItem('unPublishOfferId')))
            .subscribe(
                data => {
                    if (data.success == true) {
                        this.getPage();
                        this.hideUnPublishModal();
                    } else {
                        console.log(data.message);
                    }
                },
                err => console.error(err)
            );
    }

    showPublishModal(id: number): void {
        localStorage.setItem('publishOfferId', id.toString());
        this.publishModal.show();
    }

    hidePublishModal(): void {
        localStorage.removeItem('publishOfferId');
        this.publishModal.hide();
    }

    publish(): void {
        this.service.publish(parseInt(localStorage.getItem('publishOfferId')))
            .subscribe(
                data => {
                    if (data.success == true) {
                        this.getPage();
                        this.hidePublishModal();
                    } else {
                        console.log(data.message);
                    }
                },
                err => console.error(err)
            );
    }

    delete(): void {
        this.service.delete(parseInt(localStorage.getItem('deleteRowId')))
            .subscribe(
                data => {
                    if (data.success == true) {
                        this.getPage();
                        this.hideDeleteModal();
                    } else {
                        console.log(data.message);
                    }
                },
                err => console.error(err)
            );
    }

    changeSearch(val: string): void {
        this.getPage(1, val);
    }

    changePerPage(perPage: number): void {
        this.getPage(parseInt(localStorage.getItem('currentPage')), localStorage.getItem('searchText'), perPage);
    }

    editModel(id: number): void {
        this.router.navigate([this.listUrl + '/' + id]);
    }

}