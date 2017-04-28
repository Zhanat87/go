import {Component, OnInit} from '@angular/core';
import {Title} from "@angular/platform-browser";

@Component({
    selector: 'page-not-found',
    templateUrl: './404.html',
})
export class PageNotFoundComponent implements OnInit {

    public constructor(private titleService: Title) {

    }

    ngOnInit() {
        this.setTitle('404 page not found');
    }

    public setTitle( newTitle: string) {
        this.titleService.setTitle(newTitle); // or document.title
    }

}