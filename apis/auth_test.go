package apis

import (
	"testing"
	"net/http"
)

func TestAuth(t *testing.T) {
	router.Post("/auth/sign-in", SignIn("secret"))
	runAPITests(t, []apiTestCase{
		{"t1 - successful login", "POST", "/auth/sign-in", `{"email":"test@test.com", "password":"pass"}`, http.StatusOK, ""},
		{"t2 - unsuccessful login", "POST", "/auth/sign-in", `{"email":"demo@demo.com", "password":"bad"}`, http.StatusUnauthorized, ""},
	})
}