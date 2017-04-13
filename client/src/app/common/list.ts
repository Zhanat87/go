import { OnInit } from '@angular/core';

export abstract class CommonList implements OnInit {
    public mode = 'Observable';
    
    public service;
    
    public data;
    
    public errorMessage: string;

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
