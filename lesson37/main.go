package main

import (
	"log"
	"github.com/Go11Group/Intizor/lesson37/handler"
    "github.com/Go11Group/Intizor/lesson37/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("No connect database ", err.Error())
	}

	u := postgres.NewUserRepo(db)

	r := handler.NewHandler(handler.Handler{User: u})

	err = r.Run(":8023")

	if err != nil {
		log.Fatal("No running gin ", err.Error())
	}
	
}