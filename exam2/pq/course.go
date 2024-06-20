package pq

import (
	"database/sql"
	"log"

	"github.com/Go11Group/Intizor/exam2/models"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{db}
}

func (db *CourseRepo) CreateCourse(req *models.Course) error {

	_, err := db.DB.Exec(
		`insert into courses(title, description) values($1, $2)`, req.Title, req.Description)
	return err
}
func (db *CourseRepo) ById(id string) (*models.Course, error) { //id boyicha courseni oladi

	course := models.Course{}

	query :=
		`SELECT 
	 	course_id,
		title,
		description,
		create_at,
		update_at,
		deleted_at
	FROM
		courses
	WHERE 
		id = $1 `

	row := db.DB.QueryRow(query, id) //sorovni bajarish uchun queryrow dan foydalanib qatorni qaytaradi
	err := row.Scan(                 //qaytadigan malumotlarni scan qilib oqib oladi va models.Course course ozgaruvchiga saqlab qoyadi.
		&course.CourseID,
		&course.Title,
		&course.Description,
		&course.CreatedAt,
		&course.UpdatedAt,
		&course.DeletedAt,
	)

	if err != nil {
		log.Println("error get course by id: ", err)
	}

	return &course, nil
}

func (db *CourseRepo) GetCoursesByUser(userID string) (*[]models.Course, error) {
    var courses []models.Course

    query := `
        SELECT 
            course_id,
            title,
            description,
            create_at,
            update_at,
            deleted_at
        FROM
            courses
        WHERE 
            user_id = $1 AND deleted_at = 0`

    rows, err := db.DB.Query(query, userID)
    if err != nil {
        log.Println("error querying courses by user id: ", err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var course models.Course
        err := rows.Scan(
            &course.CourseID,
            &course.Title,
            &course.Description,
            &course.CreatedAt,
            &course.UpdatedAt,
            &course.DeletedAt,
        )
        if err != nil {
            log.Println("error scanning course: ", err)
            return nil, err
        }
        courses = append(courses, course)
    }

    if err = rows.Err(); err != nil {
        log.Println("rows error: ", err)
        return nil, err
    }

    return &courses, nil
}

func (db *CourseRepo) GetAllCourses(title string) (*[]models.Course, error) {
	resp := []models.Course{}

	query := `SELECT * FROM courses WHERE deleted_at = 0`
	var args []interface{}

	if title != "" {
		query += " AND title = ?"
		args = append(args, title)
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close() 

	for rows.Next() {
		course := models.Course{}
		err := rows.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
		if err != nil {
			return nil, err
		}
		resp = append(resp, course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (db *CourseRepo) UpdateCourse(id string, newtitle string) error {

	_, err := db.DB.Exec(
		`update courses set title=$1 where course_id=$2 and deleted_at=0`, newtitle, id,
	)
	return err
}

func (db *CourseRepo) DeleteCourse(id string) error {

	_, err := db.DB.Exec(
		`update courses set deleted_at = date_part('epoch', current_timestamp)::INT where courses_id=$1`, id,
	)
	return err
}

//date_part('epoch', current_timestamp)::INT -> 13245678908765
