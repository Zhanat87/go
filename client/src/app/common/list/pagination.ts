import { OnInit, OnDestroy, ViewChild } from '@angular/core';
import { Observable } from 'rxjs/Rx';
import {CommonPaginationFooterComponent} from "../components/pagination/footer/footer";
import {Router} from "@angular/router";
import {GlobalState} from "../../global.state";
import {BreadCrumb} from "../entities/breadCrumb";
import {ServerError} from "../entities/serverError";
import {LocalStorageService} from "angular-2-local-storage";

export abstract class CommonListPagination implements OnInit, OnDestroy {
    public active = false;

    public search = true;

    // @todo: need interface type here
    public service;

    public data;

    public perPage: number;

    public currentPage: number = 1;

    public lastPage: number = 1;

    public total: number;

    public search_params = {};

    public errorMessage: string;

    public router: Router;

    public listUrl: string;

    public title: string;

    protected _state: GlobalState;

    protected localStorageService: LocalStorageService;

    @ViewChild(CommonPaginationFooterComponent) vc:CommonPaginationFooterComponent;

    ngOnInit() {
        this.getPage(1);
        this.setBreadCrumbs();
    }

    ngOnDestroy() {
        localStorage.removeItem('perPage');
        localStorage.removeItem('searchText');
        localStorage.removeItem('currentPage');
        localStorage.removeItem('lastPage');
        localStorage.removeItem('totalPage');

        localStorage.removeItem('breadCrumbs');
    }

    getPage(page?: number, search?: string, perPage?: number) {
        this.active = false;

        page = page ? page : (localStorage.getItem('currentPage') ? parseInt(localStorage.getItem('currentPage')) : 1);
        perPage = perPage ? perPage : (localStorage.getItem('perPage') ? parseInt(localStorage.getItem('perPage')) : 15);
        search = search ? (search != ' ' ? search : '') : (localStorage.getItem('searchText') ? localStorage.getItem('searchText') : '');

        let params = {
            'page': page,
            'search': search,
            'perPage': perPage,
        };
        localStorage.setItem('currentPage', page.toString());
        localStorage.setItem('searchText', search);
        localStorage.setItem('perPage', perPage.toString());

        let observableBatch = [
            this.service.paginate(params)
                .do((res: any) => {
                    let data = res.json(); // as PaginationResponse
                    this.total = data.total_count;
                    this.perPage = data.per_page;
                    this.currentPage = page; // data.page
                    this.lastPage = data.page_count;

                    localStorage.setItem('lastPage', data.page_count.toString());
                    localStorage.setItem('totalPage', data.total_count.toString());

                    this.vc.updatePageLinks();
                })
                .map((res: any) => res.json().items)
        ];

        Observable.forkJoin(
            observableBatch
        ).subscribe(
            data => {
                this.data = data[0];
            },
            err => {
                console.error(err);
                // let error = err as ServerError;
                // if (error.status_code == 401) {
                //     console.log('ServerError', error.message);
                //     this.logout();
                // }
            }
        );
    }

    protected setBreadCrumbs(): void {
        let breadCrumbs = [];
        breadCrumbs.push(new BreadCrumb(this.title));
        localStorage.setItem('breadCrumbs', JSON.stringify(breadCrumbs));
        this._state.notifyChanged('breadCrumbs');
    }

    private logout(): void {
        localStorage.clear();
        localStorage.setItem('referrer', window.location.pathname);
        this.router.navigate(['/login']);
    }

}