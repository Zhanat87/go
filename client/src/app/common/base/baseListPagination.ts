import {CommonListPagination} from "../list/pagination";
import { ViewChild } from '@angular/core';
import {CommonDeleteModalComponent} from "../components/modals/delete/delete";

export abstract class BaseListPagination extends CommonListPagination {

    @ViewChild(CommonDeleteModalComponent) dm: CommonDeleteModalComponent;

    ngOnDestroy() {
        super.ngOnDestroy();

        this.localStorageService.remove('deleteRowId');
    }

    showDeleteModal(id: number): void {
        this.dm.showDeleteModal(id);
    }

    delete(): void {
        this.service.delete(this.localStorageService.get('deleteRowId'))
            .subscribe(
                data => {
                    if (data.id) {
                        this.getPage();
                        this.dm.hideDeleteModal();
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

    /**
     * @link https://angular.io/docs/ts/latest/api/common/index/NgFor-directive.html
     * @link https://netbasal.com/angular-2-improve-performance-with-trackby-cc147b5104e5
     *
     * @param row
     * @returns mixed
     */
    track(row: any): any {
        return row ? row.id : undefined;
    }

    trackByFn(index, item) {
        return index; // or item.id
    }

}
