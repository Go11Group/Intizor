package main

import (
	"fmt"
	"handler"
	packages "storage/postgres"
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