import { OnInit, OnDestroy } from '@angular/core';

import { Router, ActivatedRoute } from '@angular/router';
import {GlobalState} from "../global.state";
import {BreadCrumb} from "./entities/breadCrumb";
import {OnAfterSave} from "./interfaces/form_lifecycle_hooks";
import { LocalStorageService } from 'angular-2-local-storage';

export abstract class CommonForm implements OnInit, OnDestroy, OnAfterSave {
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

    protected localStorageService: LocalStorageService;

    public title: string;

    public statuses = {
        1 : 'active',
        2 : 'suspended',
        3 : 'deleted',
    };

    public weekDays = [
        'monday',
        'tuesday',
        'wednesday',
        'thursday',
        'friday',
        'saturday',
        'sunday',
    ];

    public months = [
        'january',
        'february',
        'march',
        'april',
        'may',
        'june',
        'july',
        'august',
        'september',
        'october',
        'november',
        'december',
    ];

    ngOnDestroy(): void {
        this.localStorageService.remove('breadCrumbs');
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
        this.localStorageService.set('breadCrumbs', JSON.stringify(breadCrumbs));
        this._state.notifyChanged('breadCrumbs');
    }

    abstract getBreadCrumbTitle(): string;

    onAfterSave(): void {
        this.redirectList();
    }

}
