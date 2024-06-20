package handler

import (
	"net/http"

	"github.com/Go11Group/Intizor/exam2/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateEnrollment(c *gin.Context) {

	req := models.Enrollment{}
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.EP.CreateEnrollment(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok :)"})
}

func (h *Handler) GetEnrolledUsersByCourse(c *gin.Context) {
    courseID := c.Param("course_id")

    users, err := h.EP.GetEnrolledUsersByCourse(courseID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}

func (h *Handler) GetEnrollmentById(c *gin.Context) {
	id := c.Param("id")
	enroll, err := h.EP.GetEnrollmentById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enroll)
}
func (h *Handler) GetAllEnrollments(c *gin.Context) {
    enrollmentDate := c.Query("enrollment_date")

    resp, err := h.EP.GetAllEnrollments(enrollmentDate)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp)
}


func (h *Handler) DeleteEnrollment(c *gin.Context) {

	id := c.Query("id")

	err := h.EP.DeleteEnrollment(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok :)"})
}
