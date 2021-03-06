package usersDb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

const (
	username  = "FILL_ME"
	password  = "FILL_ME"
	host_port = "127.0.0.1:3306"
	schema    = "users_db"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host_port,
		schema,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName) // Don't use := , use = , https://github.com/go-sql-driver/mysql/issues/150
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
