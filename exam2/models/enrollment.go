package models

import (
	"time"
)

type Enrollment struct {
	EnrollmentID   string    `db:"enrollment_id" json:"enrollment_id"`
	UserID         string    `db:"user_id" json:"user_id"`
	CourseID       string    `db:"course_id" json:"course_id"`
	EnrollmentDate time.Time `db:"enrollment_date" json:"enrollment_date"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt      int64     `db:"deleted_at" json:"deleted_at"`
}
