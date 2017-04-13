export class BreadCrumb {

    title: string;
    link?: string;

    public constructor(title: string, link?: string) {
        this.title = title;
        this.link  = link;
    }

}