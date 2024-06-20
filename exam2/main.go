package main

import (
	"log"

	"github.com/Go11Group/Intizor/exam2/api"
	"github.com/Go11Group/Intizor/exam2/api/handler"
	"github.com/Go11Group/Intizor/exam2/pq"
)

func main() {

	db, err := pq.ConnectDB()
	if err != nil {
		log.Println("error while opening database", err)
	}
	defer db.Close()

	h := handler.NewHandler(
		pq.NewUserRepo(db),
		pq.NewCourseRepo(db),
		pq.NewLessonRepo(db),
		pq.NewEnrollmentRepo(db),
	)

	r := api.NewGin(h)
	err = r.Run(":8081")
	if err != nil {
		log.Println("error while starting server")
	}
}
