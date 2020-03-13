package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB ...
type DB struct {
	SQL *sql.DB
	// Mgo *mgo.database
}

// DBConn ...
var dbConn = &DB{}

// ConnectSQL ...
func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	// dbSource := fmt.Sprintf(
	// 	"root:%s@tcp(%s:%s)/%s?charset=utf8",
	// 	pass,
	// 	host,
	// 	port,
	// 	dbname,
	// )
	dbSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		uname,
		pass,
		dbname)

	d, err := sql.Open("postgres", dbSource)
	if err != nil {
		panic(err)
	}
	dbConn.SQL = d

	err = d.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	return dbConn, err
}
