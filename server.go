package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var variables string
	for _, e := range os.Environ() {
		variables += e + "\r\n"
	}
	dbPort, issetPort := os.LookupEnv("POSTGRESQL_PORT")
	if issetPort {
		variables += "db port: " + dbPort + "\r\n"
		dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("POSTGRESQL_ENV_POSTGRES_USER"),
			os.Getenv("POSTGRESQL_ENV_POSTGRES_PASSWORD"), os.Getenv("POSTGRESQL_PORT_5432_TCP_ADDR"),
			os.Getenv("POSTGRESQL_PORT_5432_TCP_PORT"), os.Getenv("POSTGRESQL_ENV_POSTGRES_DB"))
		variables += "dns: " + dns + "\r\n"
	}
	_, issetPort2 := os.LookupEnv("POSTGRESQL_PORT2")
	if issetPort2 == false {
		variables += "db port2: not isset\r\n"
	}

	fmt.Fprintf(w, variables)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
/*
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
	db, err := dbx.MustOpen("postgres", app.Config.GetDSN())
	if err != nil {
		panic(err)
	}
	db.LogFunc = logger.Infof

	// wire up API routing
	http.Handle("/", buildRouter(logger, db))

	// start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(logger *logrus.Logger, db *dbx.DB) *routing.Router {
	router := routing.New()

	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		c.Abort()  // skip all other middlewares/handlers
		return c.Write("OK " + app.Version)
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

	// wire up more resource APIs here

	return router
}
*/
/*
STACK_POSTGRES_PORT=tcp://172.17.0.2:5432
STACK_POSTGRES_ENV_PG_MAJOR=9.6
HOSTNAME=c6093c27814a
STACK_POSTGRES_PORT_5432_TCP=tcp://172.17.0.2:5432
STACK_POSTGRES_NAME=/golang/stack-postgres
STACK_POSTGRES_ENV_POSTGRES_DB=stack
HOME=/root
STACK_POSTGRES_ENV_PGDATA=/var/lib/postgresql/data
GOLANG_DOWNLOAD_SHA256=47fda42e46b4c3ec93fa5d4d4cc6a748aa3f9411a2a2b7e08e3a6d80d753ec8b
STACK_POSTGRES_ENV_PG_VERSION=9.6.2-1.pgdg80+1
PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
STACK_POSTGRES_ENV_LANG=en_US.utf8
GOPATH=/go
STACK_POSTGRES_PORT_5432_TCP_ADDR=172.17.0.2
STACK_POSTGRES_ENV_POSTGRES_PASSWORD=stack
STACK_POSTGRES_ENV_GOSU_VERSION=1.7
PWD=/go
STACK_POSTGRES_PORT_5432_TCP_PORT=5432
STACK_POSTGRES_PORT_5432_TCP_PROTO=tcp
STACK_POSTGRES_ENV_POSTGRES_USER=stack
GOLANG_DOWNLOAD_URL=https://golang.org/dl/go1.7.4.linux-amd64.tar.gz
GOLANG_VERSION=1.7.4
 */
