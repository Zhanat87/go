/*
 http://stackoverflow.com/questions/36528824/pass-environment-variables-to-an-angular2-app
 http://stackoverflow.com/questions/41694053/setup-the-environment-variables-of-angular-2-project
 https://github.com/AngularClass/angular2-webpack-starter/wiki/How-to-pass-environment-variables%3F
 */
export class Environment {

    public static API_ENDPOINT = window.location.href.substr(0, 16) == 'http://localhost' ?
        'http://localhost:8080/' : 'http://zhanat.site:8080/';

    public static IS_LOCAL = window.location.href.substr(0, 16) == 'http://localhost';

    public static SOCKET_URL = window.location.href.substr(0, 16) == 'http://localhost' ?
        'http://localhost:5000/' : 'http://zhanat.site:5000/';

}