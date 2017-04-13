import {Component, EventEmitter, Input, Output} from '@angular/core';

export interface Page {
    label: string;
    value: any;
}

/**
 * This directive is what powers all pagination controls components, including the default one.
 * It exposes an API which is hooked up to the PaginationService to keep the PaginatePipe in sync
 * with the pagination controls.
 */
@Component({
    selector: 'pagination-footer',
    templateUrl: './footer.html',
})
export class CommonPaginationFooterComponent {
    @Input() id: string;
    @Input() maxSize: number = 7;
    @Output() pageChange: EventEmitter<number> = new EventEmitter<number>();
    pages: Page[] = [];

    /**
     * Go to the previous page
     */
    previous() {
        this.setCurrent(this.getCurrent() - 1);
    }

    /**
     * Go to the next page
     */
    next() {
        this.setCurrent(this.getCurrent() + 1);
    }
    /**
     * Returns true if current page is first page
     */
    isFirstPage(): boolean {
        return this.getCurrent() == 1;
    }
    /**
     * Returns true if current page is last page
     */
    isLastPage(): boolean {
        return this.getCurrent() == this.getLastPage();
    }

    /**
     * Get the current page number.
     */
    getCurrent(): number {
        return parseInt(localStorage.getItem('currentPage'));
    }
    /**
     * Returns the last page number
     */
    getLastPage(): number {
        return parseInt(localStorage.getItem('lastPage'));
    }

    ngOnInit() {
        if (this.id === undefined) {
            this.id = 'paginationId';
        }
        this.updatePageLinks();
    }

    ngOnChanges(changes: any) {
        this.updatePageLinks();
    }

    /**
     * Set the current page number.
     */
    setCurrent(page: number) {
        this.pageChange.emit(page);
    }

    /**
     * Updates the page links and checks that the current page is valid. Should run whenever the
     * PaginationService.change stream emits a value matching the current ID, or when any of the
     * input values changes.
     */
    updatePageLinks() {
        this.pages = this.createPageArray(this.getCurrent(), parseInt(localStorage.getItem('perPage')), parseInt(localStorage.getItem('totalPage')), this.maxSize);
    }

    /**
     * Returns an array of Page objects to use in the pagination controls.
     */
    private createPageArray(currentPage: number, itemsPerPage: number, totalItems: number, paginationRange: number): Page[] {
        // paginationRange could be a string if passed from attribute, so cast to number.
        paginationRange = +paginationRange;
        let pages = [];
        const totalPages = Math.ceil(totalItems / itemsPerPage);
        const halfWay = Math.ceil(paginationRange / 2);

        const isStart = currentPage <= halfWay;
        const isEnd = totalPages - halfWay < currentPage;
        const isMiddle = !isStart && !isEnd;

        let ellipsesNeeded = paginationRange < totalPages;
        let i = 1;

        while (i <= totalPages && i <= paginationRange) {
            let label;
            let pageNumber = this.calculatePageNumber(i, currentPage, paginationRange, totalPages);
            let openingEllipsesNeeded = (i === 2 && (isMiddle || isEnd));
            let closingEllipsesNeeded = (i === paginationRange - 1 && (isMiddle || isStart));
            if (ellipsesNeeded && (openingEllipsesNeeded || closingEllipsesNeeded)) {
                label = '...';
            } else {
                label = pageNumber;
            }
            pages.push({
                label: label,
                value: pageNumber
            });
            i ++;
        }
        return pages;
    }

    /**
     * Given the position in the sequence of pagination links [i],
     * figure out what page number corresponds to that position.
     */
    private calculatePageNumber(i: number, currentPage: number, paginationRange: number, totalPages: number) {
        let halfWay = Math.ceil(paginationRange / 2);
        if (i === paginationRange) {
            return totalPages;
        } else if (i === 1) {
            return i;
        } else if (paginationRange < totalPages) {
            if (totalPages - halfWay < currentPage) {
                return totalPages - paginationRange + i;
            } else if (halfWay < currentPage) {
                return currentPage - halfWay + i;
            } else {
                return i;
            }
        } else {
            return i;
        }
    }
}