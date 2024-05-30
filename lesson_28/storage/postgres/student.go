package postgres

import (
    "database/sql"
    "github.com/Go11Group/at_lesson/lesson28/model"
)

type StudentRepo struct {
    Db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
    return &StudentRepo{Db: db}
}

func (u *StudentRepo) GetAllStudents() ([]model.User, error) {
    rows, err := u.Db.Query(`SELECT s.id, s.name, s.age, s.gender, c.name as course
                             FROM student s
                             LEFT JOIN course c ON c.id = s.course_id`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []model.User
    for rows.Next() {
        var user model.User
        err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func (u *StudentRepo) GetByID(id string) (model.User, error) {
    var user model.User

    err := u.Db.QueryRow(`SELECT s.id, s.name, s.age, s.gender, c.name as course
                          FROM student s
                          LEFT JOIN course c ON c.id = s.course_id 
                          WHERE s.id = $1`, id).
        Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
    if err != nil {
        if err == sql.ErrNoRows {
            return model.User{}, nil
        }
        return model.User{}, err
    }

    return user, nil
}

func (u *StudentRepo) Create(user model.User) (model.User, error) {
    err := u.Db.QueryRow(`INSERT INTO student (name, age, gender, course_id)
                          VALUES ($1, $2, $3, $4)
                          RETURNING id, name, age, gender, course_id`,
        user.Name, user.Age, user.Gender, user.Course_id).
        Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course_id)
    if err != nil {
        return model.User{}, err
    }

    return user, nil
}

func (u *StudentRepo) Update(user model.User) error {
    _, err := u.Db.Exec(`UPDATE student
                         SET name = $1, age = $2, gender = $3, course_id = $4
                         WHERE id = $5`,
        user.Name, user.Age, user.Gender, user.Course_id, user.ID)
    return err
}

func (u *StudentRepo) Delete(id string) error {
    _, err := u.Db.Exec(`DELETE FROM student
                         WHERE id = $1`, id)
    return err
}
