package dbconnect

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB(DBDriver, DBSource string) (db *sql.DB) {
	db, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}
