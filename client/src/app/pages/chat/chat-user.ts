export class ChatUser {

    userId: number;
    username: string;
    avatar: string;

    constructor(userId: number, username: string, avatar: string) {
        this.userId = userId;
        this.username = username;
        this.avatar = avatar;
    }

}