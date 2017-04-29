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
	_ "github.com/lib/pq"
	"github.com/Zhanat87/go/apis"
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/daos"
	"github.com/Zhanat87/go/errors"
	"github.com/Zhanat87/go/services"
	"os"
	"database/sql"
)

func main() {
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
		for _, e := range os.Environ() {
			variables += e + "\r\n"
		}
		_, err := sql.Open("postgres", dsn)
		if err != nil {
			variables += err.Error() + "\r\n"
		} else {
			variables += "success connected\r\n"
		}
		return c.Write(variables + "\r\n" + dsn + "\r\ndeploy\r\naot compilation\r\nwebhook22\r\n")
	})

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

	rg := router.Group("/v1")

	rg.Post("/auth", apis.Auth(app.Config.JWTSigningKey))
	rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
		SigningMethod: app.Config.JWTSigningMethod,
		TokenHandler:  apis.JWTHandler,
	}))

	artistDAO := daos.NewArtistDAO()
	apis.ServeArtistResource(rg, services.NewArtistService(artistDAO))

	albumDAO := daos.NewAlbumDAO()
	apis.ServeAlbumResource(rg, services.NewAlbumService(albumDAO))

	userDAO := daos.NewUserDAO()
	apis.ServeUserResource(rg, services.NewUserService(userDAO))

	// wire up more resource APIs here

	return router
}
