export class ChatMessage {

    userId: number;
    username: string;
    message: string;
    avatar: string;
    time: string;

    constructor(userId: number, username: string, message: string, avatar: string) {
        this.userId = userId;
        this.username = username;
        this.message = message;
        this.avatar = avatar;
        let date = new Date();
        this.time = date.getHours() + ':' + date.getMinutes() + ':' + date.getSeconds();
    }

}