import { OnInit, OnDestroy } from '@angular/core';

import { Router, ActivatedRoute } from '@angular/router';
import {GlobalState} from "../global.state";
import {BreadCrumb} from "./entities/breadCrumb";

export abstract class CommonForm implements OnInit, OnDestroy {
    public listUrl: string;

    public service;

    public model;
    
    public router: Router;

    public route: ActivatedRoute;
    
    public errorMessage: string;

    public submitted = false;

    public active = false;

    public editMode = false;

    protected _state: GlobalState;

    public title: string;

    public activations = {
        active : 'main.active',
        suspended : 'main.suspended',
        deleted : 'main.deleted',
    };

    public weekDays = [
        'contract.monday',
        'contract.tuesday',
        'contract.wednesday',
        'contract.thursday',
        'contract.friday',
        'contract.saturday',
        'contract.sunday'
    ];

    public months = [
        'contract.january',
        'contract.february',
        'contract.march',
        'contract.april',
        'contract.may',
        'contract.june',
        'contract.july',
        'contract.august',
        'contract.september',
        'contract.october',
        'contract.november',
        'contract.december'
    ];

    ngOnDestroy(): void {
        localStorage.removeItem('breadCrumbs');
    }

    ngOnInit() {
        this.route.params.subscribe(params => {
            let id = +params['id'];

            if (id) {
                this.editMode = true;

                this.service.find(id)
                    .subscribe(
                        data => {
                            this.active = true;
                            this.model = data;
                        },
                        error => this.errorMessage = <any>error);
            } else {
                this.active = true;
            }
        });
    }

    onSubmit() {
        this.active = false;
        this.submitted = true;

        let values = this.getValues();

        if (this.editMode == true) {
            this.service.update(values, this.model.id)
                .subscribe(
                    data => {
                        this.model = data;
                        this.redirectList();
                    },
                    error => this.errorMessage = <any>error);
        } else {
            this.service.create(values)
                .subscribe(
                    data =>  {
                        this.model = data;
                        this.redirectList();
                    },
                    error => this.errorMessage = <any>error);
        }
    }

    delete() {
        if (this.editMode == true) {
            this.service.delete(this.model.id)
                .subscribe(
                    data => this.redirectList(),
                    error => this.errorMessage = <any>error);
        }
    }

    // TODO: Перенести в сервис
    redirectList() {
        // let url = this.service.parseUrl(this.listUrl);
        // this.router.navigate([url]);
        this.router.navigate([this.listUrl]);
    }

    getValues() {
        return this.model;
    }

    protected setBreadCrumbs(): void {
        let breadCrumbs = [];
        breadCrumbs.push(new BreadCrumb(this.title, this.listUrl));
        breadCrumbs.push(new BreadCrumb(this.getBreadCrumbTitle()));
        localStorage.setItem('breadCrumbs', JSON.stringify(breadCrumbs));
        this._state.notifyChanged('breadCrumbs');
    }

    abstract getBreadCrumbTitle(): string;

}
