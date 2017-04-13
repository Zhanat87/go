package responses

import "github.com/Zhanat87/go/models"

type SignInSuccessResponse struct {
	APISuccess
	Data SignInData `json:"data"`
}

type SignInData struct  {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func MakeSignInSuccessResponse(token string, identity models.Identity) SignInSuccessResponse {
	return SignInSuccessResponse{
		APISuccess: APISuccess{Status: 200, Message: "ok"},
		Data: SignInData{
			Token: token,
			User:  models.User{ID: identity.GetID(), Name: identity.GetName(), Email: identity.GetEmail()},
		},
	}
}