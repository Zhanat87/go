import { OnInit } from '@angular/core';
import { LocalStorageService } from 'angular-2-local-storage';

export abstract class CommonList implements OnInit {
    public mode = 'Observable';
    
    public service;
    
    public data;
    
    public errorMessage: string;

    protected localStorageService: LocalStorageService;

    ngOnInit() { 
        this.getAll(); 
    }

    getAll() {
        this.service.all()
            .subscribe(
                data => this.data = data,
                error =>  this.errorMessage = <any>error);
    }
}
