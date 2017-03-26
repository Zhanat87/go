package responses

import "github.com/Zhanat87/go/models"

type SignInSuccessResponse struct {
	APISuccess
	Data SignInData `json:"data"`
}

type SignInData struct  {
	Token    string `json:"token"`
	Username string `json:"username"`
}

func MakeSignInSuccessResponse(token string, identity models.Identity) SignInSuccessResponse {
	return &SignInSuccessResponse{
		APISuccess: APISuccess{Status: 200, Message: "ok"},
		Data: SignInData{
			Token: token,
			Username: identity.GetName(),
		},
	}
}