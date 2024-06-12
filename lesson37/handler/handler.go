package handler

import (
    "github.com/Go11Group/Intizor/lesson37/postgres" //bu shaxsiy paket, unda UserRepo aniqlangan
	"github.com/gin-gonic/gin"// bu veb-server yaratish funcsiyalarini ta'minlovchi Gin veb ramkasi

	_ "github.com/lib/pq" // bu blok kerakli pacetlarni import qiladi
)

type Handler struct {
	User *postgres.UserRepo
} // bu struct postgres.UserRepo tipidagi User maydonini o'z ichiga olgan Handler turini belgilaydi

func NewHandler(handler Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/users2", handler.GetUsers)

	return r
}


