package main

import (
	"github.com/Go11Group/Intizor/Lesson43/api_gateway_service/api"
	"github.com/Go11Group/Intizor/Lesson43/api_gateway_service/api/handler"
	"log"
)

func main() {
	router := api.Router(handler.Handler{})

	log.Fatal(router.Run(":8080"))
}