package handler

import (
	"net/http"

	"github.com/Go11Group/Intizor/exam2/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {

	req := models.User{}
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.UP.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok :)"})
}
func (h *Handler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := h.UP.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) SearchUsers(c *gin.Context) {
    email := c.Query("email")

    users, err := h.UP.Search(email)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
    name := c.Query("name")

    resp, err := h.UP.GetAllUsers(name)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (h *Handler)UpdatePassword(c *gin.Context) {

	id := c.Query("id")
	newPassword := c.Query("password")

	err := h.UP.UpdatePassword(id, newPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok :)"})
}

func (h *Handler) DeleteUser(c *gin.Context) {

	id := c.Query("id")

	err := h.UP.DeleteUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok :)"})
}
