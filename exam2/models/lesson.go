package models

import(
	"time"
)


type Lesson struct {
    LessonID   string `db:"lesson_id" json:"lesson_id"`
    CourseID   string `db:"course_id" json:"course_id"`
    Title      string    `db:"title" json:"title"`
    Content    string    `db:"content" json:"content"`
    CreatedAt  time.Time `db:"created_at" json:"created_at"`
    UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
    DeletedAt  int64     `db:"deleted_at" json:"deleted_at"`
}