import { Injectable } from '@angular/core';
import {Router, CanActivate, CanActivateChild, ActivatedRouteSnapshot, RouterStateSnapshot} from '@angular/router';
import {Response, Http, Headers, RequestOptions} from "@angular/http";
import {Observable} from "rxjs";
import {tokenNotExpired, JwtHelper} from "angular2-jwt";
import {Environment} from "../environment";
import {RefreshToken} from "../entities/refreshToken";

/**
 * https://angular.io/docs/ts/latest/api/router/index/CanActivate-interface.html
 */
@Injectable()
export class AuthGuard implements CanActivate, CanActivateChild {

    constructor(private router: Router, private http: Http) {}

    canActivate(): Observable<boolean> | Promise<boolean> | boolean {
        return this.checkAccess();
    }

    canActivateChild(childRoute: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<boolean> | Promise<boolean> | boolean {
        return this.checkAccess();
    }

    checkAccess(needCheck?: boolean): Observable<boolean> | Promise<boolean> | boolean {
        if (needCheck === false) {
            return true;
        }
        if (localStorage.getItem('id_token')) {
            let date = new Date;
            let jwtHelper = new JwtHelper;
            if (jwtHelper.getTokenExpirationDate(localStorage.getItem('id_token')).getTime() < date.getTime() + 60000) {
                this.refreshToken();
            } else {
                return tokenNotExpired(null, localStorage.getItem('id_token'));
            }
        } else {
            localStorage.setItem('referrer', window.location.hash.substr(1));
            this.router.navigate(['/login']);
            return false;
        }
    }

    refreshToken() {
        this.refreshTokenQuery()
            .subscribe(
                data => {
                    let res = data as RefreshToken;
                    if (res.message == 'token_refreshed') {
                        localStorage.setItem('id_token', res.data.token);
                        this.checkAccess(false);
                    }
                },
                error => {
                    /*
                     {"message":"Token has expired and can no longer be refreshed","status_code":500}
                     {"message":"Token has expired","status_code":401}
                     */
                    console.log('refresh token error', error);
                    this.logout();
                },
                () => {
                    console.log('refresh token done');
                },
            );
    }

    refreshTokenQuery() : Observable<RefreshToken> {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Authorization', 'Bearer ' + localStorage.getItem('id_token'));
        let options = new RequestOptions({ headers: headers });

        return this.http
            .patch(Environment.API_ENDPOINT + 'auth/refresh', null, options)
            .map((res:Response) => res.json())
            .catch((error:any) => Observable.throw(error || 'Server error'));
    }

    private logout(): void {
        localStorage.clear();
        localStorage.setItem('referrer', window.location.hash.substr(1));
        this.router.navigate(['/login']);
    }

}