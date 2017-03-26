package responses

type SignInResponse struct {
	APISuccess
	Data SignInData `json:"data"`
}

type SignInData struct  {
	Token    string `json:"token"`
	Username string `json:"username"`
}