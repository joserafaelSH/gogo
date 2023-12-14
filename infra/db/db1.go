package db

import (
	"database/sql"
	"fmt"

	"github.com/joserafaelsh/syncrhonize-database/infra/config"
	_ "github.com/lib/pq"
)

func ConnectDB1() *sql.DB {

	postgresqlDbInfo := config.CreteUrlConnectionDb1()

	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}
