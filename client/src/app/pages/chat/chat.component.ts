import {Component, OnInit, OnDestroy, ViewEncapsulation} from '@angular/core';
import {GlobalState} from "../../global.state";
import { LocalStorageService } from 'angular-2-local-storage';
import {BreadCrumb} from "../../common/entities/breadCrumb";
import * as io from "socket.io-client";
import {Environment} from "../../common/environment";
import {User} from "../../modules/user/user";
import {ChatMessage} from "./chat-message";
import {ChatUser} from "./chat-user";

@Component({
    selector: 'chat',
    styleUrls: ['./chat.scss'],
    templateUrl: './chat.html',
    /*
    angular 2 css not working after add new element
    @link https://github.com/angular/angular/issues/7845
    @link http://stackoverflow.com/questions/36265026/angular-2-innerhtml-styling
     */
    encapsulation: ViewEncapsulation.None,
})
export class ChatComponent implements OnInit, OnDestroy {

    public title = 'Chat';

    public errorMessage: string;

    public active = false;

    currentUser: User;

    private socket: any;

    public disabledSendMessage: boolean = true;

    constructor(private _state: GlobalState,
                private localStorageService: LocalStorageService) {
    }

    ngOnInit(): void {
        this.setData();
    }

    ngOnDestroy(): void {
        if (this.socket) {
            this.socket.emit('forceDisconnect');
            this.socket.emit('chatUserLogout', this.getUserId());
            this.socket.disconnect();
        }
    }

    setData(): void {
        this.active = true;

        this.setBreadCrumbs();
        this.setCurrentUser();
        this.initSocket();
    }

    initSocket(): void {
        this.socket = io(Environment.SOCKET_URL, {transports: ['websocket']});

        let self = this;
        // listen for messages
        this.socket.on('chatMessage', function(msg: ChatMessage) {
            self.addMessage(msg, self.getUserId() == msg.userId);
        });
        this.socket.on('chatUsers', function(user: ChatUser) {
            self.addUser(user);
        });
        this.socket.on('chatUserLogout', function(userId: number) {
            self.deleteUser(userId);
        });

        let socket = this.socket;
        this.socket.on('connect', function () {
            console.log('socket connected');
            // send something
            socket.emit('chatUsers', self.createChatUser(), function(result: string) {
                console.log('chatUsers ack: ', result);
            });
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

    sendMessage(): void {
        let $messageInput = jQuery(document.getElementById('messageInput'));
        let msg = $messageInput.val();
        $messageInput.val('');

        this.socket.emit('chatMessage', this.createChatMessage(msg), function(result: string) {
            console.log('chatMessage ack: ', result);
        });
    }

    private getUserName(): string {
        return this.currentUser.full_name ? this.currentUser.full_name :
            (this.currentUser.username ? this.currentUser.username : this.currentUser.email);
    }

    private getAvatarUrl(): string {
        return this.currentUser.avatar_string ? (this.currentUser.avatar_string.substr(0, 4) == 'http' ?
            this.currentUser.avatar_string : Environment.API_ENDPOINT + 'static/users/avatars/' +
            this.currentUser.avatar_string) : Environment.API_ENDPOINT + 'static/img/default-avatar.jpg';
    }

    private getUserId(): number {
        return this.currentUser.id;
    }

    onMessageInputChange(event): void {
        let target = event.currentTarget || event.target || event.srcElement;
        this.disabledSendMessage = target.value == '';
    }

    addMessage(msg: ChatMessage, self?: boolean): void {
        let chatMsg = `
            <li class="${self ? 'left' : 'right'} clearfix">
                <span class="chat-img pull-${self ? 'left' : 'right'}">
                    <img src="${msg.avatar}" />
                </span>
                <div class="chat-body clearfix">
                    <div class="header">
                        <strong class="primary-font">${self ? 'me' : msg.username}</strong>
                        <small class="pull-right text-muted"><i class="fa fa-clock-o"></i> ${msg.time}</small>
                    </div>
                    <p>
                        ${msg.message}
                    </p>
                </div>
            </li>
            `;
        jQuery(document.getElementById('chatUl')).append(chatMsg);
    }

    private createChatMessage(msg: string): ChatMessage {
        return new ChatMessage(this.getUserId(), this.getUserName(), msg, this.getAvatarUrl());
    }

    private createChatUser(): ChatUser {
        return new ChatUser(this.getUserId(), this.getUserName(), this.getAvatarUrl());
    }

    private deleteUser(userId: number): void {
        jQuery('li[id=userId' + userId + ']').remove();
    }

    addUser(msg: ChatUser): void {
        let v = '';
        if (this.getUserId() == msg.userId) {
            v = `
            <li class="active bounceInDown">
                <a class="clearfix">
                    <img src="${msg.avatar}" alt="" class="img-circle">
                    <div class="friend-name">
                        <strong>me</strong>
                    </div>
                </a>
            </li>
            `;
        } else {
            v = `
            <li id="userId${msg.userId}">
                <a class="clearfix">
                    <img src="${msg.avatar}" alt="" class="img-circle">
                    <div class="friend-name">
                        <strong>${msg.username}</strong>
                    </div>
                </a>
            </li>
            `;
        }
        jQuery(document.getElementById('chatUsersUl')).append(v);
    }

}