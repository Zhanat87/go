package apis

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/errors"
	"github.com/Zhanat87/go/models"
	"github.com/Zhanat87/go/responses"
	"golang.org/x/crypto/bcrypt"
)

type Credential struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth(signingKey string) routing.Handler {
	return func(c *routing.Context) error {
		var credential Credential
		if err := c.Read(&credential); err != nil {
			return errors.Unauthorized(err.Error())
		}

		identity := authenticate(credential)
		if identity == nil {
			return errors.Unauthorized("invalid credential")
		}

		token, err := auth.NewJWT(jwt.MapClaims{
			"id":   identity.GetID(),
			"name": identity.GetName(),
			"email": identity.GetEmail(),
			"exp":  time.Now().Add(time.Hour * 72).Unix(),
		}, signingKey)
		if err != nil {
			return errors.Unauthorized(err.Error())
		}

		return c.Write(responses.MakeSignInSuccessResponse(token, identity))
	}
}

func authenticate(c Credential) models.Identity {
	if (c.Username == "demo" || c.Email == "user@demo.com") && validatePassword(c.Password) {
		return &models.User{ID: "100", Name: "demo", , Email: "user@demo.com"}
	}
	return nil
}

func validatePassword(password string) bool {
	// "pass" hash
	hashedPassword := []byte("$2a$10$YOGE3lBg7SXbhEa8kr8B3OBFimlWLrytjad8VquOFWBYIVY1UP.xa")
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return false
	}
	return true
}

func JWTHandler(c *routing.Context, j *jwt.Token) error {
	userID := j.Claims.(jwt.MapClaims)["id"].(string)
	app.GetRequestScope(c).SetUserID(userID)
	return nil
}
