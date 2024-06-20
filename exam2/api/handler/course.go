package handler

import (
	"net/http"

	"github.com/Go11Group/Intizor/exam2/models"
	"github.com/gin-gonic/gin"
)


func(h *Handler) CreateCourse(c *gin.Context) {
	
	req := models.Course{}
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.CP.CreateCourse(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"ok :)"})
} 
func (h *Handler) GetCourseById(c *gin.Context) {
	id := c.Param("id")
	course, err := h.CP.ById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}

func (h *Handler) GetCoursesByUser(c *gin.Context) {
    userID := c.Param("user_id")

    courses, err := h.CP.GetCoursesByUser(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, courses)
}


func (h *Handler) GetAllCourses(c *gin.Context) {
    title := c.Query("title")

    resp, err := h.CP.GetAllCourses(title)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp)
}


func (h *Handler) UpdateCoursePassword(c *gin.Context) {

	id := c.Query("id")
	newtitle := c.Query("title")

	err := h.CP.UpdateCourse(id, newtitle)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} 
	c.JSON(http.StatusOK, gin.H{"message":"ok :)"})
}


func (h *Handler) DeleteCourse(c *gin.Context) {

	id := c.Query("id")

	err := h.CP.DeleteCourse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} 
	c.JSON(http.StatusOK, gin.H{"message":"ok :)"})
}

