package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"
	"github.com/go-ozzo/ozzo-routing/file"
	_ "github.com/lib/pq"
	"github.com/Zhanat87/go/apis"
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/daos"
	"github.com/Zhanat87/go/errors"
	"github.com/Zhanat87/go/services"
	"os"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/joho/godotenv"
	"github.com/Zhanat87/go/helpers"
	errs "errors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		helpers.FailOnError(err, "Error loading .env file", true)
	}

	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// load error messages
	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	// create the logger
	logger := logrus.New()

	// connect to the database
	//db, err := dbx.MustOpen("postgres", app.Config.GetDSN())
	db, err := dbx.Open("postgres", app.Config.GetDSN())
	if err != nil {
		// docker compose can't start and docker image can't build, need first timeout for connect
		//time.Sleep(20 * time.Second)
		//db, err = dbx.MustOpen("postgres", app.Config.GetDSN())
		//if err != nil {
			panic(err)
		//}
	}
	db.LogFunc = logger.Infof

	// wire up API routing
	http.Handle("/", buildRouter(logger, db, app.Config.GetDSN()))

	// start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(logger *logrus.Logger, db *dbx.DB, dsn string) *routing.Router {
	router := routing.New()

	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		c.Abort()  // skip all other middlewares/handlers
		return c.Write("OK " + app.Version)
	})

	router.To("GET,HEAD", "/test", func(c *routing.Context) error {
		c.Abort()  // skip all other middlewares/handlers
		var variables string
		if c.Get("secret_key").(string) == os.Getenv("SECRET_KEY") {
			for _, e := range os.Environ() {
				variables += e + "\r\n"
			}
		} else {
			variables += c.Get("secret_key").(string) + "\r\n"
		}
		_, err := sql.Open("postgres", dsn)
		if err != nil {
			variables += err.Error() + "\r\n"
		} else {
			variables += "success connected\r\n"
		}

		return c.Write(variables + "\r\n" + dsn + "\r\ndeploy\r\naot compilation works now\r\nwebhook\r\n" +
			"avatar crop upload\r\n")
	})

	if helpers.IsDocker() == false {
		router.To("GET", "/telegram/<msg>", func(c *routing.Context) error {
			helpers.LogError(errs.New(c.Param("msg")))
			return c.Write(fmt.Sprintf("send message to telegram success: %s", c.Param("msg")))
		})
	}

	// serve static files
	router.Get("/static/*", file.Server(file.PathMap{
		"/": "/",
	}))

	router.Use(
		app.Init(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
		app.Transactional(db),
	)
	// note: from here start access to db
	userDAO := daos.NewUserDAO()
	router.Get("/auth/*", apis.SocialAuth(userDAO))

	rg := router.Group("/v1")

	userService := services.NewUserService(userDAO)
	rg.Post("/auth/sign-up", apis.SignUp(userService))
	rg.Post("/auth/password-reset-request", apis.PasswordResetRequest(userDAO))
	rg.Post("/auth/password-reset/<token>", apis.PasswordReset(userDAO))

	rg.Post("/auth/sign-in", apis.SignIn(app.Config.JWTSigningKey))
	// @link https://github.com/go-ozzo/ozzo-routing#handlers
	rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
		SigningMethod: app.Config.JWTSigningMethod,
		TokenHandler:  apis.JWTHandler,
	}))
	rg.Delete("/auth/sign-out", apis.SignOut())
	rg.Patch("/auth/refresh-jwt-token", apis.RefreshJWTToken(app.Config.JWTSigningKey))

	rg.Get("/restricted", func(c *routing.Context) error {
		claims := c.Get("JWT").(*jwt.Token).Claims.(jwt.MapClaims)

		unixTime := int64(claims["exp"].(float64))
		t := time.Now().Unix()
		delta := unixTime - t
	        return c.Write(fmt.Sprintf("id: %d, username: %s, email: %s, exp: %v, exp date: %v, delta: %v, ct: %v",
			claims["id"], claims["username"], claims["email"], claims["exp"], unixTime, delta, t))
	})

	rg.Get("/user-id", func(c *routing.Context) error {
		userId := app.GetRequestScope(c).UserID()

		return c.Write(fmt.Sprintf("user id: %d", userId))
	})

	artistDAO := daos.NewArtistDAO()
	apis.ServeArtistResource(rg, services.NewArtistService(artistDAO))

	albumDAO := daos.NewAlbumDAO()
	apis.ServeAlbumResource(rg, services.NewAlbumService(albumDAO))

	apis.ServeUserResource(rg, userService)

	categoryDAO := daos.NewCategoryDAO()
	apis.ServeCategoryResource(rg, services.NewCategoryService(categoryDAO))

	// partitions
	newsDAO := daos.NewNewsDAO()
	apis.ServeNewsResource(rg, services.NewNewsService(newsDAO))

	// shards
	newsShardDAO := daos.NewNewsShardDAO()
	apis.ServeNewsShardResource(rg, services.NewNewsShardService(newsShardDAO))

	// replications
	newsReplicationDAO := daos.NewNewsReplicationDAO()
	apis.ServeNewsReplicationMasterResource(rg, services.NewNewsReplicationService(newsReplicationDAO))
	apis.ServeNewsReplicationSlaveResource(rg, services.NewNewsReplicationService(newsReplicationDAO))

	return router
}
