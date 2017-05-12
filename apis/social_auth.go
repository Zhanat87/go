package apis

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/google"
	"net/http"
	"strings"
	"os"
	"fmt"
	"time"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/Zhanat87/go/models"
	"errors"
	"github.com/Zhanat87/go/helpers"
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/daos"
	"database/sql"
)

func SocialAuth(userDAO *daos.UserDAO) routing.Handler {
	return func(c *routing.Context) error {
		segs := strings.Split(c.Request.URL.Path, "/")
		action := segs[2]
		providerType := segs[3]
		uuid := segs[4]

		// setup gomniauth
		gomniauth.SetSecurityKey(os.Getenv("GOMNIAUTH_SECURITY_KEY"))
		gomniauth.WithProviders(
			github.New(os.Getenv("GITHUB_CLIENT_ID"), os.Getenv("GITHUB_CLIENT_SECRET"),
				os.Getenv("API_BASE_URL") + "auth/callback/github/" + uuid),
			google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"),
				os.Getenv("API_BASE_URL") + "auth/callback/google/" + uuid),
			facebook.New(os.Getenv("FACEBOOK_APP_ID"), os.Getenv("FACEBOOK_SECRET_KEY"),
				os.Getenv("API_BASE_URL") + "auth/callback/facebook/" + uuid),
		)

		var response error

		switch action {
		case "login":

			provider, err := gomniauth.Provider(providerType)
			if err != nil {
				return socialAuthError("Error when trying to get provider", providerType, err)
			}

			loginUrl, err := provider.GetBeginAuthURL(nil, nil)
			if err != nil {
				return socialAuthError("Error when trying to GetBeginAuthURL for", providerType, err)
			}
			c.Response.Header()["Location"] = []string{loginUrl}
			c.Response.WriteHeader(http.StatusTemporaryRedirect)

		case "callback":

			provider, err := gomniauth.Provider(providerType)
			if err != nil {
				return socialAuthError("Error when trying to get provider", providerType, err)
			}

			// get the credentials
			creds, err := provider.CompleteAuth(objx.MustFromURLQuery(c.Request.URL.RawQuery))
			if err != nil {
				return socialAuthError("Error when trying to complete auth for: " + c.Request.URL.Query().Get("code"), providerType, err)
			}

			user, err := provider.GetUser(creds)
			if err != nil {
				return socialAuthError("Error when trying to get user from", providerType, err)
			}

			/*
			if user exist with this provider and provider_id, so simple update his data from social net,
			if not exist create new user,
			if exist email, show message that user with this email exist
			 */
			var errorVar error
			rs := app.GetRequestScope(c)
			model, err := userDAO.FindByEmail(rs, user.Email())
			if err != nil {
				if err.Error() == "sql: no rows in result set" {
					model.Email = user.Email()
					model.Username = user.Nickname()
					model.FullName = user.Name()
					if len(user.AvatarURL()) > 0 {
						model.AvatarString = models.JsonNullString{sql.NullString{String: user.AvatarURL(), Valid:true}}
					}
					model.Provider = models.JsonNullString{sql.NullString{String: providerType, Valid:true}}
					model.ProviderId = models.JsonNullString{sql.NullString{String: user.IDForProvider(providerType), Valid:true}}
					model.Status = 1;
					err = userDAO.Create(rs, model)
					if err != nil {
						return err
					}
				} else {
					return err
				}
			} else {
				if model.Provider.String == providerType && model.ProviderId.String == user.IDForProvider(providerType) {
					model.Username = user.Nickname()
					model.FullName = user.Name()
					if len(user.AvatarURL()) > 0 {
						model.AvatarString = models.JsonNullString{sql.NullString{String: user.AvatarURL(), Valid:true}}
					}
					err = userDAO.Update(rs, model.Id, model)
					if err != nil {
						return err
					}
				} else {
					errorVar = errors.New(fmt.Sprintf("User with this email %s exist. " +
						"If you don't remember you password, you can restore it.", user.Email()))
				}
			}

			var token string
			if errorVar == nil {
				token, err = createToken(model, os.Getenv("JWT_SIGNING_KEY"))
				if err != nil {
					return err
				}
			}

			sendSocialAuthMessage(model, uuid, token, errorVar)

			response = c.Write([]byte(fmt.Sprintf("name: %s\r\nemail: %s\r\nusername: %s\r\n" +
				"avatar url: %s\r\nid for provider: %s\r\ndata: %v\r\nuuid: %s\r\n",
				user.Name(), user.Email(), user.Nickname(), user.AvatarURL(),
				user.IDForProvider(providerType), user.Data(), uuid)))
		default:
			c.Response.WriteHeader(http.StatusNotFound)
			response = c.Write([]byte(fmt.Sprintf("Auth action %s not supported", action)))
		}

		return response
	}
}

/*
~/go/src/github.com/graarh/golang-socketio/examples$ go run server.go
~/go/src/github.com/graarh/golang-socketio/examples$ go run client.go

@link https://socket.io/docs/client-api/
@link https://github.com/graarh/golang-socketio
@link https://github.com/googollee/go-socket.io
 */
type SocialAuthMessage struct {
	Uuid  string       `json:"uuid"`
	User  *models.User `json:"user"`
	Token string       `json:"token"`
	Error error        `json:"error"`
}

func sendSocialAuthMessage(user *models.User, uuid, token string, errorVar error) (error, bool) {
	// connect to server, you can use your own transport settings
	conn, err := gosocketio.Dial(
		gosocketio.GetUrl("localhost", 5000, false),
		transport.GetDefaultWebsocketTransport(),
	)
	defer conn.Close()

	if err != nil {
		socketError := errors.New(fmt.Sprintln("Error when connect to socket io server", err))
		helpers.LogError(socketError)
		return socketError, false
	}

	conn.Emit("socialAuth", SocialAuthMessage{Uuid: uuid, Token: token, User: user, Error: errorVar})
	time.Sleep(1 * time.Second)

	return nil, true
}

func socialAuthError(providerType, msg string, err error) error {
	return errors.New(fmt.Sprintln(msg, providerType, "-", err))
}