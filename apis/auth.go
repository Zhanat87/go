package apis

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/errors"
	"github.com/Zhanat87/go/responses"
	"github.com/Zhanat87/go/daos"
	"github.com/Zhanat87/go/models"
	"strings"
	"github.com/Zhanat87/go/db"
	"github.com/Zhanat87/go/helpers"
	golang_errors "errors"
)

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
https://godoc.org/golang.org/x/oauth2/jwt
https://jwt.io/
 */
func SignIn(signingKey string) routing.Handler {
	return func(c *routing.Context) error {
		var credential Credential
		err := c.Read(&credential)
		if err != nil {
			return errors.Unauthorized(err.Error())
		}

		rs := app.GetRequestScope(c)
		var user *models.User
		if strings.Contains(credential.Email, "@") {
			user, err = daos.NewUserDAO().FindByEmail(rs, credential.Email)
			if err != nil {
				return errors.Unauthorized("user with this email not found")
			}
		} else {
			user, err = daos.NewUserDAO().FindByUsername(rs, credential.Email)
			if err != nil {
				return errors.Unauthorized("user with this username not found")
			}
		}
		if !user.ValidatePassword(credential.Password) {
			return errors.Unauthorized("password not valid")
		}

		token, err := auth.NewJWT(jwt.MapClaims{
			"id":       user.GetId(),
			"username": user.GetUsername(),
			"email":    user.GetEmail(),
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		}, signingKey)
		if err != nil {
			return errors.Unauthorized(err.Error())
		}

		return c.Write(responses.MakeSignInSuccessResponse(token, user))
	}
}

func JWTHandler(c *routing.Context, j *jwt.Token) error {
	if ok, err := helpers.CheckJWTTokenIsValid(c); ok != true {
		return golang_errors.New(err.Error())
	}
	// @link http://stackoverflow.com/questions/18041334/convert-interface-to-int-in-golang
	userID := int(j.Claims.(jwt.MapClaims)["id"].(float64))
	app.GetRequestScope(c).SetUserID(userID)
	// src/github.com/go-ozzo/ozzo-routing/auth/handlers.go:208
	c.Set("JWT", j)
	return nil
}

func SignOut() routing.Handler {
	return func(c *routing.Context) error {
		claims := c.Get("JWT").(*jwt.Token).Claims.(jwt.MapClaims)
		unixTime := int64(claims["exp"].(float64))
		t := time.Now().Unix()
		delta := unixTime - t

		redis := db.NewRedis()
		err := redis.Set(helpers.GetJWTToken(c), true, time.Second * time.Duration(delta)).Err()
		if err != nil {
			panic(err)
		}

		return c.Write(responses.APISuccess{Status: 200, Message: "token_invalidated"})
	}
}