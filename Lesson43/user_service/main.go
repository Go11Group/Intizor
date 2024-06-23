package main

import (
	"log"
	"github.com/Go11Group/Intizor/Lesson43/user_service/api"
	"github.com/Go11Group/Intizor/Lesson43/user_service/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := api.Router(db)
	err = server.Run(":8082")
	log.Println("Running on :8082")
	if err != nil {
		panic(err)
	}

}
