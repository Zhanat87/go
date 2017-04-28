import {User} from "../../modules/user/user";

export class LoginResponse {
    status: string;
    message: string;
    data: LoginResponseData;
}

class LoginResponseData {
    token: string;
    user: User;
}
