package apis

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/errors"
	"github.com/Zhanat87/go/responses"
	//golang_errors "errors"
	//"strconv"
	"github.com/Zhanat87/go/daos"
	"github.com/Zhanat87/go/models"
)

type Credential struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth(signingKey string) routing.Handler {
	return func(c *routing.Context) error {
		var credential Credential
		err := c.Read(&credential)
		if err != nil {
			return errors.Unauthorized(err.Error())
		}

		rs := app.GetRequestScope(c)
		var user *models.User
		if len(credential.Email) > 0 {
			user, err = daos.NewUserDAO().FindByEmail(rs, credential.Email)
			if err != nil {
				return errors.Unauthorized("user with this email not found")
			}
		} else {
			user, err = daos.NewUserDAO().FindByUsername(rs, credential.Username)
			if err != nil {
				return errors.Unauthorized("user with this username not found")
			}
		}
		if !user.ValidatePassword(credential.Password) {
			return errors.Unauthorized("password not valid")
		}

		token, err := auth.NewJWT(jwt.MapClaims{
			"id":   user.GetId(),
			"username": user.GetUsername(),
			"email": user.GetEmail(),
			"exp":  time.Now().Add(time.Hour * 72).Unix(),
		}, signingKey)
		if err != nil {
			return errors.Unauthorized(err.Error())
		}

		return c.Write(responses.MakeSignInSuccessResponse(token, user))
	}
}

func JWTHandler(c *routing.Context, j *jwt.Token) error {
	app.GetRequestScope(c).SetUserID(123)
	return nil
	//if userID, ok := j.Claims.(jwt.MapClaims)["id"].(string); ok {
	//	userID, err := strconv.Atoi(userID)
	//	if err != nil {
	//		return err
	//	}
	//	app.GetRequestScope(c).SetUserID(userID)
	//	return nil
	//}
	//return golang_errors.New("JWTHandler: bad user id")
}
