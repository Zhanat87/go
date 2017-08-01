package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/Zhanat87/go/helpers"
	"github.com/Zhanat87/go/app"
	//"os"
)

var Connection *sql.DB

func init() {
	//Dsn := "host=" + os.Getenv("DB_HOST") +
	//	" port=" + os.Getenv("DB_PORT") +
	//	" user=" + os.Getenv("DB_USER") +
	//	" password=" + os.Getenv("DB_PASSWORD") +
	//	" dbname=" + os.Getenv("DB_DATABASE") +
	//	" sslmode=disable"
	Dsn := app.Config.GetDSN()
	var err error
	Connection, err = sql.Open("postgres", Dsn)
	helpers.FailOnError(err, "error postgres connection", false)
}
