package postgres

import (
	"database/sql"
	"github.com/Go11Group/at_lesson/lesson28/model"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
	return &CourseRepo{DB}
}

func (c *CourseRepo) Create(course *model.Course) error {
	query := `INSERT INTO course (name, field) VALUES ($1, $2) RETURNING id`
	err := c.DB.QueryRow(query, course.Name, course.Field).Scan(&course.Id)
	return err
}

func (c *CourseRepo) GetByID(id string) (*model.Course, error) {
	query := `SELECT id, name, field FROM course WHERE id = $1`
	course := &model.Course{}
	err := c.DB.QueryRow(query, id).Scan(&course.Id, &course.Name, &course.Field)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (c *CourseRepo) Update(course *model.Course) error {
	query := `UPDATE course SET name = $1, field = $2 WHERE id = $3`
	_, err := c.DB.Exec(query, course.Name, course.Field, course.Id)
	return err
}

func (c *CourseRepo) Delete(id string) error {
	query := `DELETE FROM course WHERE id = $1`
	_, err := c.DB.Exec(query, id)
	return err
}

func (c *CourseRepo) GetAllCourses() ([]*model.Course, error) {
	query := `SELECT id, name, field FROM course`
	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*model.Course
	for rows.Next() {
		course := &model.Course{}
		if err := rows.Scan(&course.Id, &course.Name, &course.Field); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
