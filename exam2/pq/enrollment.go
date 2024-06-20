package pq

import (
	"database/sql"
	"log"

	"github.com/Go11Group/Intizor/exam2/models"
)

type EnrollmentRepo struct {
	DB *sql.DB
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{db}
}

func (db *EnrollmentRepo) CreateEnrollment(req *models.Enrollment) error {

	_, err := db.DB.Exec(
		`insert into enrollments(enrollment_date) values($1)`, req.EnrollmentDate)
	return err
}
func (db *EnrollmentRepo) GetEnrollmentById(id string) (*models.Enrollment, error) { //id boyicha enrollmentni oladi

	enrollment := models.Enrollment{}

	query :=
		`SELECT 
	 	enrollment_id,
		user_id,
		course_id,
		enrollment_date,
		create_at,
		update_at,
		deleted_at
	FROM
		enrollments
	WHERE 
		id = $1 `

	row := db.DB.QueryRow(query, id) //sorovni bajarish uchun queryrow dan foydalanib qatorni qaytaradi
	err := row.Scan(                 //qaytadigan malumotlarni scan qilib oqib oladi va models.Enrollment enrolllment ozgaruvchiga saqlab qoyadi.
		&enrollment.EnrollmentID,
		&enrollment.UserID,
		&enrollment.CourseID,
		&enrollment.EnrollmentDate,
		&enrollment.CreatedAt,
		&enrollment.UpdatedAt,
		&enrollment.DeletedAt,
	)

	if err != nil {
		log.Println("error get enrollment by id: ", err)
	}

	return &enrollment, nil
}

func (db *EnrollmentRepo) GetEnrolledUsersByCourse(courseID string) (*[]models.User, error) {
    var users []models.User

    query := `
        SELECT 
            u.user_id,
            u.name,
            u.email,
            u.birthday,
            u.password,
            u.created_at,
            u.updated_at,
            u.deleted_at
        FROM
            users u
        JOIN 
            enrollments e ON u.user_id = e.user_id
        WHERE 
            e.course_id = $1 AND u.deleted_at = 0`

    rows, err := db.DB.Query(query, courseID)
    if err != nil {
        log.Println("error querying users by course id: ", err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var user models.User
        err := rows.Scan(
            &user.UserID,
            &user.Name,
            &user.Email,
            &user.Birthday,
            &user.Password,
            &user.CreatedAt,
            &user.UpdatedAt,
            &user.DeletedAt,
        )
        if err != nil {
            log.Println("error scanning user: ", err)
            return nil, err
        }
        users = append(users, user)
    }

    if err = rows.Err(); err != nil {
        log.Println("rows error: ", err)
        return nil, err
    }

    return &users, nil
}


func (db *EnrollmentRepo) GetAllEnrollments(enrollmentDate string) (*[]models.Enrollment, error) {
	resp := []models.Enrollment{}

	query := `SELECT * FROM enrollments WHERE deleted_at = 0`
	var args []interface{}

	if enrollmentDate != "" {
		query += " AND enrollment_date = ?"
		args = append(args, enrollmentDate)
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close() 

	for rows.Next() {
		enrollment := models.Enrollment{}
		err := rows.Scan(&enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
		if err != nil {
			return nil, err
		}
		resp = append(resp, enrollment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (db *EnrollmentRepo) DeleteEnrollment(id string) error {

	_, err := db.DB.Exec(
		`update enrollments set deleted_at = date_part('epoch', current_timestamp)::INT where enrollment_id=$1`, id,
	)
	return err
}

//date_part('epoch', current_timestamp)::INT -> 13245678908765
