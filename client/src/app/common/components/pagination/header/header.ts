import {Component, Input, Output, EventEmitter} from '@angular/core';

@Component({
    selector: 'pagination-header',
    templateUrl: './header.html',
})
export class CommonPaginationHeaderComponent {

    @Output() perPageChange: EventEmitter<number> = new EventEmitter<number>();
    @Output() searchChange: EventEmitter<string> = new EventEmitter<string>();

    @Input() canCreate: boolean = false;

    public perPages = [15, 50, 100];

    public searchInput: HTMLInputElement;

    onChangePerPage(perPage: any): void {
        this.perPageChange.emit(parseInt(perPage.target.value));
    }

    goSearch(val: string): void {
        this.searchChange.emit(val ? val : ' ');
    }

    get createUrl(): string {
        return window.location.hash + '/create';
    }

}