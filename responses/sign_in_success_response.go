package responses

type SignInResponse struct {
	APISuccess
	APISuccess.Data {
		Token    string `json:"token"`
		Username string `json:"username"`
	}
}
