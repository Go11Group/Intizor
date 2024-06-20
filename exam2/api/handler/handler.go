package handler

import (
	"github.com/Go11Group/Intizor/exam2/pq"
)
type Handler struct {
	UP  *pq.UserRepo
	CP  *pq.CourseRepo
	LP *pq.LessonRepo
	EP  *pq.EnrollmentRepo

}

func NewHandler(up *pq.UserRepo, cp *pq.CourseRepo, lp *pq.LessonRepo, ep *pq.EnrollmentRepo) *Handler {
	return &Handler{UP: up, CP: cp, LP: lp, EP: ep}
}
