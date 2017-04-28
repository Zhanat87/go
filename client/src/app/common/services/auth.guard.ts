import { Injectable } from '@angular/core';
import {Router, CanActivate, CanActivateChild, ActivatedRouteSnapshot, RouterStateSnapshot} from '@angular/router';
import {Response, Http, Headers, RequestOptions} from "@angular/http";
import {Observable} from "rxjs";
import {tokenNotExpired, JwtHelper} from "angular2-jwt";
import {Environment} from "../environment";
import {RefreshToken} from "../entities/refreshToken";
import { LocalStorageService } from 'angular-2-local-storage';

/**
 * https://angular.io/docs/ts/latest/api/router/index/CanActivate-interface.html
 */
@Injectable()
export class AuthGuard implements CanActivate, CanActivateChild {

    constructor(private router: Router, private http: Http, protected localStorageService: LocalStorageService) {}

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
        if (this.localStorageService.get('id_token')) {
            let date = new Date;
            let jwtHelper = new JwtHelper;
            if (jwtHelper.getTokenExpirationDate(this.localStorageService.get<string>('id_token')).getTime() < date.getTime() + 600000) {
                this.refreshToken();
            } else {
                return tokenNotExpired(null, this.localStorageService.get<string>('id_token'));
            }
        } else {
            this.localStorageService.set('referrer', window.location.pathname);
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
                        this.localStorageService.set('id_token', res.data.token);
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
        headers.append('Authorization', 'Bearer ' + this.localStorageService.get('id_token'));
        let options = new RequestOptions({ headers: headers });

        return this.http
            .patch(Environment.API_ENDPOINT + 'auth/refresh', null, options)
            .map((res:Response) => res.json())
            .catch((error:any) => Observable.throw(error || 'Server error'));
    }

    private logout(): void {
        this.localStorageService.clearAll();
        this.localStorageService.set('referrer', window.location.pathname);
        this.router.navigate(['/login']);
    }

}