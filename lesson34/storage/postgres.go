package packages

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "postgres"
	password = "root"
)

func ConnectDB() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, errpackage main

		import (
			"fmt"
			"mymod/handler"
			packages "mymod/storage/postgres"
		)
		
		type Users struct {
			Id       string
			UserName string
			Email    string
			Password string
		}
		
		
		func main() {
		
			db, err := packages.ConnectDB()
			if err != nil {
				fmt.Println(err)
				return
			}
			RepoUser := packages.RepoNewUser{Db: db}
			RepoProduct := packages.RepoNewProducts{Db: db}
			server := handler.NewHandler(RepoUser, RepoProduct)
		
			server.ListenAndServe()
		
		}
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}