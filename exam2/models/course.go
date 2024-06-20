package models

import(
	"time"
)

type Course struct {
    CourseID    string 	  `db:"course_id" json:"course_id"`
    Title       string    `db:"title" json:"title"`
    Description string    `db:"description" json:"description"`
    CreatedAt   time.Time `db:"created_at" json:"created_at"`
    UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
    DeletedAt   int64     `db:"deleted_at" json:"deleted_at"`
}