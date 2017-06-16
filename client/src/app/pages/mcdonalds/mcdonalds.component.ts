import {Component, OnInit} from '@angular/core';
import {McdonaldsItem} from "./mcdonaldsItem";
import {McdonaldsService} from "./mcdonalds.service";

@Component({
    selector: 'mcdonalds',
    styleUrls: ['./mcdonalds.scss'],
    templateUrl: './mcdonalds.html',
    providers: [
        McdonaldsService,
    ],
})
export class McdonaldsComponent implements OnInit {

    public title = 'Mcdonalds menu';

    public mcdonaldsItems: McdonaldsItem[];

    constructor(private service: McdonaldsService) {
    }

    ngOnInit(): void {
        this.service.getMcdonaldsMenu()
            .subscribe(
                data => {
                    this.mcdonaldsItems = data as McdonaldsItem[];
                },
                error => {
                    console.log('getMcdonaldsMenu error', error);
                },
            );
    }

}