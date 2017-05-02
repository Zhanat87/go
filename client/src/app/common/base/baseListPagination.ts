import {CommonListPagination} from "../list/pagination";
import {ModalDirective} from 'ngx-bootstrap/modal';
import { ViewChild } from '@angular/core';

export abstract class BaseListPagination extends CommonListPagination {

    @ViewChild('deleteModal') deleteModal: ModalDirective;
    @ViewChild('unPublishModal') unPublishModal: ModalDirective;
    @ViewChild('publishModal') publishModal: ModalDirective;

    ngOnDestroy() {
        super.ngOnDestroy();

        this.localStorageService.remove('deleteRowId');
        this.localStorageService.remove('unPublishOfferId');
        this.localStorageService.remove('publishOfferId');
    }

    showDeleteModal(id: number): void {
        this.localStorageService.set('deleteRowId', id);
        this.deleteModal.show();
    }

    hideDeleteModal(): void {
        this.localStorageService.remove('deleteRowId');
        this.deleteModal.hide();
    }

    showUnPublishModal(id: number): void {
        this.localStorageService.set('unPublishOfferId', id);
        this.unPublishModal.show();
    }

    hideUnPublishModal(): void {
        this.localStorageService.remove('unPublishOfferId');
        this.unPublishModal.hide();
    }

    unPublish(): void {
        this.service.unPublish(this.localStorageService.get('unPublishOfferId'))
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
        this.localStorageService.set('publishOfferId', id);
        this.publishModal.show();
    }

    hidePublishModal(): void {
        this.localStorageService.remove('publishOfferId');
        this.publishModal.hide();
    }

    publish(): void {
        this.service.publish(this.localStorageService.get('publishOfferId'))
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
        this.service.delete(this.localStorageService.get('deleteRowId'))
            .subscribe(
                data => {
                    if (data.id) {
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
        this.getPage(this.localStorageService.get<number>('currentPage'), this.localStorageService.get<string>('searchText'), perPage);
    }

    editModel(id: number): void {
        this.router.navigate([this.listUrl + '/' + id]);
    }

}
