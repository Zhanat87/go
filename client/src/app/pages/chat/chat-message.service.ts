import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';

import {CommonListService} from "../../common/services/list.service";
import {ChatMessage} from "./chat-message";
import {Observable} from "rxjs";
import { LocalStorageService } from 'angular-2-local-storage';

@Injectable()
export class ChatMessageService extends CommonListService {

    public url = 'v1/chat-messages';

    constructor (public http: AuthHttp, protected localStorageService: LocalStorageService) {
        super();
    }

    public map(data): Observable<ChatMessage> {
        return data;
    }

    public mapAll(data): Observable<ChatMessage[]> {
        return data;
    }

}