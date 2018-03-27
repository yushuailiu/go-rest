package mydatabase

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:yourpassword@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
