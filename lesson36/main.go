package main

import (
	"log"
	"github.com/Go11Group/Intizor/lesson36/handler"
	postgres "github.com/Go11Group/Intizor/lesson36/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Error open database", err.Error())
	}
	defer db.Close()

	u := postgres.NewUserRepo(db)

	r := handler.NewHandler(handler.Handler{User: u})

	r.Run(":8080")
}