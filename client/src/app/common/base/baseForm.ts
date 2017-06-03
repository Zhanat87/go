import { CommonForm } from '../form';
import {tryParseJSON} from "../utils";

import {AfterViewInit} from "@angular/core";
import {User} from "../../modules/user/user";

export abstract class BaseForm extends CommonForm implements AfterViewInit
{

    public errors = [];
    public serverError: string;

    public development: boolean = true;

    public id: number;

    public blocked: boolean;

    public ready: boolean = false;

    /**
     * usage:
     * (click)="isNumber($event)"
     *
     * @param event
     */
    remove(event: MouseEvent): void {
        let target = event.srcElement;
        if (confirm(target.id)) {
            this.service.delete(this.model.id)
                .subscribe(
                    data => this.redirectList(),
                    error => this.errorMessage = <any>error);
        }
    }

    /**
     * @link http://stackoverflow.com/questions/35966965/what-typescript-type-is-angular2-event
     * all events extends from UIEvent and Event
     * usage:
     * (blur)="isNumber($event)"
     *
     * @param $event
     */
    setDefaultNumber($event: FocusEvent): void {
        var target = <HTMLInputElement>$event.srcElement;
        target.value = target.value == '' ? '1' : target.value;
    }

    /**
     * @link https://www.cambiaresearch.com/articles/15/javascript-char-codes-key-codes
     * allow all digits: from 48 to 57, decimal point(46)
     * delete and backspace allow by default
     * comma (44)
     * usage:
     * (keypress)="floatNumber($event)"
     *
     * @param $event
     */
    floatNumber($event: KeyboardEvent) {
        let charCode = ($event.which) ? $event.which : $event.keyCode;
        if (charCode != 46 && (charCode < 48 || charCode > 57)) {
            this.returnFalse($event);
        }
    }

    /**
     * @link https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/keyCode
     * usage:
     * (keypress)="isNumber($event)"
     *
     * @param $event
     */
    isNumber($event: KeyboardEvent): boolean {
        let charCode = ($event.which) ? $event.which : $event.keyCode;
        if (charCode > 31 && (charCode < 48 || charCode > 57)) {
            return this.returnFalse($event);
        }
        return true;
    }

    returnFalse($event: Event): boolean {
        $event.preventDefault();
        return false;
    }

    forkJoinDone(): void {
        if (this.development) {
            console.log('Observable.forkJoin done');
        }
    }

    forkJoinError(err): void {
        if (this.development) {
            console.log('Observable.forkJoin err', err);
        }
    }

    fillModel(): void {}

    onSubmit() {
        this.submitted = true;

        this.blocked = true;

        this.fillModel();
        if (this.development) {
            console.log('model:', this.model);
        }

        if (this.editMode == true) {
            this.service.update(this.model, this.model.id)
                .subscribe(
                    data => this.response(data),
                    error => this.serverErrors(error),
                    () => this.responseDone());
        } else {
            this.service.create(this.model)
                .subscribe(
                    data => this.response(data),
                    error => this.serverErrors(error),
                    () => this.responseDone());
        }
    }

    ngOnInit(): void {
        this.route.params.subscribe(params => {
            this.id = Number(params['id']);

            if (this.id) {
                this.initEditForm(this.id);
            } else {
                this.initCreateForm();
            }
        });
    }

    ngAfterViewInit(): void {
        this.initPlugins();
    }

    initCreateForm() {
        this.active = true;
    }

    initEditForm(id: number, editMode?: boolean) {
        this.editMode = true;

        this.service.find(id).subscribe(
            data => {
                this.active = true;
                this.model = data;
            },
            err => this.forkJoinError(err),
            () => this.forkJoinDone()
        );
    }

    response(data) {
        this.blocked = false;

        if (data.error) {
            this.reInitPlugins();

            // this.errors = tryParseJSON(data.message, []);
            this.errors = data.notValid ? data.message : [];
            if (this.development) {
                console.log('error in response', data);
                this.serverError = this.errors.length == 0 ? data.message : null;
            }
        } else if (data.success) {
            this.onAfterSave();
        } else {
            this.onAfterSave(data);
            if (this.development) {
                console.log('response', data);
            }
        }
    }

    // {"error_code":"INVALID_DATA","message":"There is some problem with the data you submitted.
    // See \"details\" for more information.","details":[{"field":"password",
    // "error":"the length must be between 4 and 100"}]}
    serverErrors(errors) {
        this.blocked = false;

        this.errors = errors;
    }

    responseDone(): void {
        this.blocked = false;
    }

    /**
     * @link http://stackoverflow.com/questions/35072199/correct-way-to-use-libraries-like-jquery-jqueryui-inside-angular-2-component
     * @link https://angular.io/docs/ts/latest/guide/lifecycle-hooks.html
     *
     * note:
     * call jquery scripts not work in ngAfterViewInit() and in forkJoinDone(),
     * but jquery scripts nice works in angular directive's ngOnInit() method.
     * call in component console.log($('.someClass')) return 'angular 2 jquery n.fn.init[0]' - i.e. not found
     * call in component $(document).ready(function(){  ... your selector function...  }; too not works
     * best scenario: call jquery scripts in ngAfterViewInit() and set in constructor(el: ElementRef).
     * if in form was dynamic elements, which fills in model's property like array,
     * you must delete extra elements after form was failed send, for example:
     * contract/priceGroup/price-group-form.component.ts reInitPlugins(), for it purpose
     * need create same temp variable and use it in html template.
     */
    initPlugins(): void {}

    reInitPlugins(): void {}

    getCurrentUser(): any {
        // @link https://www.typescriptlang.org/docs/handbook/generics.html
        return JSON.parse(this.localStorageService.get<string>('currentUser')) as User;
    }

}