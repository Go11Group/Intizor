package pq

import (
	"database/sql"
	"log"

	"github.com/Go11Group/Intizor/exam2/models"
)

type LessonRepo struct {
	DB *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{db}
}

func (db *LessonRepo) CreateLesson(req *models.Lesson) error {

	_, err := db.DB.Exec(
		`insert into lessons(title, content) values($1, $2)`, req.Title, req.Content)
	return err
}

func (db *LessonRepo) GetLessonById(id string) (*models.Lesson, error) { //id boyicha lessonni oladi

	lesson := models.Lesson{}

	query :=
		`SELECT 
	 	lesson_id,
		course_id,
		title,
		content,
		create_at,
		update_at,
		deleted_at
	FROM
		lesson
	WHERE 
		id = $1 `

	row := db.DB.QueryRow(query, id) //sorovni bajarish uchun queryrow dan foydalanib qatorni qaytaradi
	err := row.Scan(                   //qaytadigan malumotlarni scan qilib oqib oladi va models.Lessoni lesson ozgaruvchiga saqlab qoyadi.
		&lesson.LessonID,
		&lesson.CourseID,
		&lesson.Title,
		&lesson.Content,
		&lesson.CreatedAt,
		&lesson.UpdatedAt,
		&lesson.DeletedAt,
	)

	if err != nil {
		log.Println("error get lesson by id: ", err)
	}

	return &lesson, nil
}

func (db *LessonRepo) GetLessonsByCourse(courseID string) (*[]models.Lesson, error) {
    var lessons []models.Lesson

    query := `
        SELECT 
            lesson_id,
            course_id,
            title,
            content,
            created_at,
            updated_at,
            deleted_at
        FROM
            lessons
        WHERE 
            course_id = $1 AND deleted_at = 0`

    rows, err := db.DB.Query(query, courseID)
    if err != nil {
        log.Println("error querying lessons by course id: ", err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var lesson models.Lesson
        err := rows.Scan(
            &lesson.LessonID,
            &lesson.CourseID,
            &lesson.Title,
            &lesson.Content,
            &lesson.CreatedAt,
            &lesson.UpdatedAt,
            &lesson.DeletedAt,
        )
        if err != nil {
            log.Println("error scanning lesson: ", err)
            return nil, err
        }
        lessons = append(lessons, lesson)
    }

    if err = rows.Err(); err != nil {
        log.Println("rows error: ", err)
        return nil, err
    }

    return &lessons, nil
}

func (db *LessonRepo) GetAllLesson(title string) (*[]models.Lesson, error) {
	resp := []models.Lesson{}

	query := `SELECT * FROM lesson WHERE deleted_at = 0`
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
		lesson := models.Lesson{}
		err := rows.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
		if err != nil {
			return nil, err
		}
		resp = append(resp, lesson)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (db *LessonRepo) UpdateLesson(id,newtitle string) error {

	_, err := db.DB.Exec(
		`update users set title=$1 where lesson_id=$2 and deleted_at=0`, newtitle, id,
	)
	return err
}

func (db *LessonRepo) DeleteLesson(id string) error {

	_, err := db.DB.Exec(
		`update lesson set deleted_at = date_part('epoch', current_timestamp)::INT where lesson_id=$1`, id,
	)
	return err
}

//date_part('epoch', current_timestamp)::INT -> 13245678908765
