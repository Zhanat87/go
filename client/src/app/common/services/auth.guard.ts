import { Injectable } from '@angular/core';
import {Router, CanActivate, CanActivateChild, ActivatedRouteSnapshot, RouterStateSnapshot} from '@angular/router';
import {Response, Http, Headers, RequestOptions} from "@angular/http";
import {Observable} from "rxjs";
import {tokenNotExpired, JwtHelper} from "angular2-jwt";
import {Environment} from "../environment";
import { LocalStorageService } from 'angular-2-local-storage';
import {trim} from "../utils";
import {SuccessResponse} from "../entities/successResponse";

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

    // if set invalidate token app redirect to login page as expected
    checkAccess(needCheck?: boolean): Observable<boolean> | Promise<boolean> | boolean {
        if (needCheck === false) {
            return true;
        }
        if (this.localStorageService.get('id_token')) {
            let date = new Date;
            let jwtHelper = new JwtHelper;
            if (jwtHelper.getTokenExpirationDate(trim(this.localStorageService.get<string>('id_token'), '"')).getTime() < date.getTime() + 600000) {
                this.refreshToken();
            } else {
                return tokenNotExpired(null, trim(this.localStorageService.get<string>('id_token'), '"'));
            }
        } else {
            this.logout();
            return false;
        }
    }

    refreshToken() {
        this.refreshTokenQuery()
            .subscribe(
                data => {
                    console.log('refresh token data', data);
                    let res = data as SuccessResponse;
                    if (res.status == 200) {
                        this.localStorageService.set('id_token', res.message);
                        this.checkAccess(false);
                    }
                },
                error => {
                    /*
                     user not found
                     server error
                     token was expired
                     token can refreshed only one time
                     */
                    console.log('refresh token error', error);
                    this.logout();
                },
                () => {
                    console.log('refresh token done');
                },
            );
    }

    refreshTokenQuery() : Observable<SuccessResponse> {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Authorization', 'Bearer ' + trim(this.localStorageService.get<string>('id_token'), '"'));
        let options = new RequestOptions({ headers: headers });

        return this.http
            .patch(Environment.API_ENDPOINT + 'v1/auth/refresh-jwt-token', null, options)
            .map((res:Response) => res.json())
            .catch((error:any) => Observable.throw(error || 'Server error'));
    }

    private logout(): void {
        this.localStorageService.clearAll();
        this.localStorageService.set('referrer', window.location.pathname);
        this.router.navigate(['/login']);
    }

}