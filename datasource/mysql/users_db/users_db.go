package users_db

import (
	"database/sql"
	"fmt"
	"log"

	// "os"
	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	mysql_user_username = "mysql_user_username"
// 	mysql_user_password = "mysql_user_password"
// 	mysql_user_host     = "mysql_user_host"
// 	mysql_user_schema   = "mysql_user_schema"
// )

var (
	Client *sql.DB

	// username:=os.Getenv(mysql_user_username)
	// password:=os.Getenv(mysql_user_password)
	// host:=os.Getenv(mysql_user_host)
	// schema:=os.Getenv(mysql_user_schema)
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root",           //mysql username
		"chengsoft",      //mysql password
		"127.0.0.1:3306", //mysql localhost
		"golang_db1",     //mysql database
	)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		//panic function is end the program
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("mysqlDatabase successful configured.")
}
