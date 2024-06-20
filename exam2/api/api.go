package api

import (
	"github.com/Go11Group/Intizor/exam2/api/handler"
	"github.com/gin-gonic/gin"
)

func NewGin(h *handler.Handler) *gin.Engine {

	r := gin.Default()

	user := r.Group("/users")
	{
		user.POST("", h.CreateUser)
		user.GET("/:id", h.GetUserById)
		user.GET("/", h.GetAllUsers)
		user.GET("/search", h.SearchUsers)
		user.PUT("/:id", h.UpdatePassword)
		user.DELETE("/:id", h.DeleteUser)
	}

	course := r.Group("/courses")
	{
		course.POST("/", h.CreateCourse)
		course.GET("/:id", h.GetCourseById)
		course.GET("/", h.GetAllCourses)
		course.GET("//user/:id", h.GetCoursesByUser)
		course.PUT("/:id", h.UpdatePassword)
		course.DELETE("/:id", h.DeleteUser)
	}

	enrollment := r.Group("/enrollments")
	{
		enrollment.POST("/", h.CreateEnrollment)
		enrollment.GET("/:id", h.GetEnrollmentById)
		enrollment.GET("/", h.GetAllEnrollments)
		// enrollment.GET("/:course_id", h.GetEnrolledUsersByCourse)
		enrollment.DELETE("/:id", h.DeleteEnrollment)
	}

	lesson := r.Group("/lessons")
	{
		lesson.POST("/", h.CreateLesson)
		lesson.GET("/:id", h.GetLessonById)
		lesson.GET("/", h.GetAllLessons)
		// lesson.GET("/:course_id", h.GetLessonsByCourse)
		lesson.PUT("/:id", h.UpdateCoursePassword)
		lesson.DELETE("/:id", h.DeleteCourse)
	}

	return r

}
