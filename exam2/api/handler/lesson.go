package handler

import (
	"net/http"

	"github.com/Go11Group/Intizor/exam2/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateLesson(c *gin.Context) {

	req := models.Lesson{}
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.LP.CreateLesson(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok :)"})
}
func (h *Handler) GetLessonById(c *gin.Context) {
	id := c.Param("id")
	lesson, err := h.LP.GetLessonById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lesson)
}

func (h *Handler) GetLessonsByCourse(c *gin.Context) {
    courseID := c.Param("course_id")

    lessons, err := h.LP.GetLessonsByCourse(courseID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, lessons)
}

func (h *Handler) GetAllLessons(c *gin.Context) {
    title := c.Query("title")

    resp, err := h.LP.GetAllLesson(title)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp)
}


func (h *Handler) UpdateLesson(c *gin.Context) {

	id := c.Query("id")
	newtitle := c.Query("titlw")

	err := h.LP.UpdateLesson(id, newtitle)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok :)"})
}

func (h *Handler) DeleteLesson(c *gin.Context) {

	id := c.Query("id")

	err := h.LP.DeleteLesson(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok :)"})
}
