/*
 http://stackoverflow.com/questions/36528824/pass-environment-variables-to-an-angular2-app
 http://stackoverflow.com/questions/41694053/setup-the-environment-variables-of-angular-2-project
 https://github.com/AngularClass/angular2-webpack-starter/wiki/How-to-pass-environment-variables%3F
 */
export class Environment {
    public static API_ENDPOINT = process.env.ENV == 'production' ?
        'http://zhanat.site:8080/' : 'http://localhost:8080/';
}