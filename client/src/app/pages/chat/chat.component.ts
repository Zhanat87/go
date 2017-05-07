import {Component, OnInit} from '@angular/core';
import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';
import {BreadCrumb} from "../../common/entities/breadCrumb";
import * as io from "socket.io-client";
import {Environment} from "../../common/environment";
import {User} from "../../modules/user/user";
import {ChatMessage} from "./chat-message";
import {ChatMessageService} from "./chat-message.service";
import {Observable} from "rxjs/Observable";

@Component({
    selector: 'chat',
    styleUrls: ['./chat.scss'],
    templateUrl: './chat.html',
    providers: [
        ChatMessageService,
    ],
})
export class ChatComponent implements OnInit {

    public title = 'Chat';

    public errorMessage: string;

    public active = false;

    socket: any = null;

    currentUser: User;

    public chatMessages: ChatMessage[];

    constructor(private _state: GlobalState,
                private localStorageService: LocalStorageService,
                public service: ChatMessageService) {
    }

    ngOnInit(): void {
        this.setData();
    }

    setData(): void {
        Observable.forkJoin(
            this.service.all()
        ).subscribe(
            data => {
                this.active = true;

                this.chatMessages = data[0] as ChatMessage[];

                console.log('this.chatMessages', this.chatMessages);

                this.setBreadCrumbs();
                this.setCurrentUser();
                this.initSocket();
            },
            error => this.errorMessage = <any>error);
    }

    initSocket(): void {
        this.socket = io(Environment.SOCKET_URL);

        let currentUser = this.currentUser;
        this.socket.on('chat message', function(msg){
            let avatar = currentUser.avatar ? currentUser.avatar : Environment.SOCKET_URL + 'default-avatar.jpg';
            let chatMsg = `
            <li class="right clearfix">
                <span class="chat-img pull-right">
                    <img src="${avatar}" />
                </span>
                <div class="chat-body clearfix">
                    <div class="header">
                        <strong class="primary-font">${currentUser.username}</strong>
                        <small class="pull-right text-muted"><i class="fa fa-clock-o"></i> 13 mins ago</small>
                    </div>
                    <p>
                        ${msg}
                    </p>
                </div>
            </li>
            `;
            jQuery(document.getElementById('chatUl')).append(chatMsg);
        });
    }

    protected setBreadCrumbs(): void {
        let breadCrumbs = [];
        breadCrumbs.push(new BreadCrumb(this.title));
        this.localStorageService.set('breadCrumbs', JSON.stringify(breadCrumbs));
        this._state.notifyChanged('breadCrumbs');
    }

    setCurrentUser(): void {
        this.currentUser = JSON.parse(this.localStorageService.get<string>('currentUser')) as User;
    }

    addNewChat(event): void {

    }

}