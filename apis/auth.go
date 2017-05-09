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
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"fmt"
	"os"
	"github.com/go-ozzo/ozzo-dbx"
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

		token, err := createToken(user, signingKey)
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

		err := db.NewRedis().Set(helpers.GetJWTToken(c), true, time.Second * time.Duration(delta)).Err()
		if err != nil {
			errors.InternalServerError(err)
		}

		return c.Write(responses.APISuccess{Status: 200, Message: "token_invalidated"})
	}
}

func SignUp(service userService) routing.Handler {
	return func(c *routing.Context) error {
		var model models.User
		if err := c.Read(&model); err != nil {
			return err
		}
		model.Username = model.Email
		model.Status = 1
		_, err := service.Create(app.GetRequestScope(c), &model)
		if err != nil {
			return err
		}

		return c.Write(responses.APISuccess{Status: 200, Message: "user success registered!"})
	}
}

func RefreshJWTToken(signingKey string) routing.Handler {
	return func(c *routing.Context) error {
		client := db.NewRedis()
		token := helpers.GetJWTToken(c)
		_, err := client.Get(token + "_refresh").Result()
		if err == redis.Nil {
			userDAO := daos.NewUserDAO()
			user, err := userDAO.Get(app.GetRequestScope(c), app.GetRequestScope(c).UserID())
			if err != nil {
				return errors.NotFound("user")
			}
			token, err := createToken(user, signingKey)
			if err != nil {
				return errors.Unauthorized(err.Error())
			}

			client.Set(token + "_refresh", true, time.Hour * 24).Err()
			if err != nil {
				errors.InternalServerError(err)
			}

			return c.Write(responses.APISuccess{Status: 200, Message: token})
		} else if err != nil {
			return errors.InternalServerError(err)
		} else {
			return errors.Unauthorized("token can refreshed only one time")
		}
	}
}

func createToken(user *models.User, signingKey string) (string, error) {
	return auth.NewJWT(jwt.MapClaims{
		"id":       user.GetId(),
		"username": user.GetUsername(),
		"email":    user.GetEmail(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}, signingKey)
}

func PasswordResetRequest(userDAO *daos.UserDAO) routing.Handler {
	return func(c *routing.Context) error {
		// @link https://github.com/go-ozzo/ozzo-routing#reading-request-data
		data := &struct{
			Email string
		}{}
		if err := c.Read(&data); err != nil {
			return err
		}

		rs := app.GetRequestScope(c)
		user, err := userDAO.FindByEmail(rs, data.Email)
		if err != nil {
			return err
		}

		token := uuid.NewV4().String()
		// @link https://groups.google.com/forum/#!topic/golang-nuts/2MzMl_sff4E
		user.PasswordResetToken = &token
		msg := fmt.Sprintf("<a href='%spassword-reset/%s' target='_blank'>reset password</a>",
			os.Getenv("CLIENT_BASE_URL"), token)
		_, err = helpers.SendEmail(user.Email, "Golang app: request password reset", msg)
		if err != nil {
			return err
		}
		dbBuilder := rs.Tx()
		_, err = dbBuilder.Update("user", dbx.Params{"password_reset_token": token},
			dbx.HashExp{"id": user.Id}).Execute()
		if err != nil {
			return err
		}

		return c.Write(responses.APISuccess{Status: 200, Message: "Check your email for further instructions!"})
	}
}

func PasswordReset(userDAO *daos.UserDAO) routing.Handler {
	return func(c *routing.Context) error {
		// @link https://github.com/go-ozzo/ozzo-routing#reading-request-data
		data := &struct{
			Password string
		}{}
		if err := c.Read(&data); err != nil {
			return err
		}

		rs := app.GetRequestScope(c)
		user, err := userDAO.FindByField(rs, "password_reset_token", c.Param("token"))
		if err != nil {
			return err
		}
		hash, err := user.Hash(data.Password)
		if err != nil {
			return err
		}

		dbBuilder := rs.Tx()
		_, err = dbBuilder.Update("user", dbx.Params{"password_reset_token": nil, "password_hash": hash},
			dbx.HashExp{"id": user.Id}).Execute()
		if err != nil {
			return err
		}

		return c.Write(responses.APISuccess{Status: 200, Message: "Password success changed!"})
	}
}