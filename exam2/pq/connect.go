package pq

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	user = "postgres"
	dbname = "postgres"
	password = "root"
	port = 5432
)

func ConnectDB() (*sql.DB, error) {

	dataSourceName := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable", host, user, dbname, password, port)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}