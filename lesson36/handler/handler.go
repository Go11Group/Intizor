package handler

import (
	postgres "github.com/Go11Group/Intizor/lesson36/postgres"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	User *postgres.UserRepo
}

func NewHandler(handler Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/api/users", handler.getUsers)
	r.POST("/api/users", handler.CreateUser)
	r.GET("/api/users/:id", handler.getUser)
	r.PUT("/api/users/:id", handler.updateUser)
	r.DELETE("/api/users/:id", handler.deleteUser)

	return r
}