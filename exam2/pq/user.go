package pq

import (
	"database/sql"
	"log"
	"strings"

	"github.com/Go11Group/Intizor/exam2/models"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (db *UserRepo) CreateUser(req *models.User) error {

	_, err := db.DB.Exec(
		`insert into users(name, email, birthday, password) values($1, $2, $3, $4)`, req.Name, req.Email, req.Birthday, req.Password)
	return err
}
func (user *UserRepo) GetUserById(id string) (*models.User, error) { //id boyicha userni oladi

	users := models.User{}

	query :=
		`SELECT 
	 	user_id,
		name,
		email,
		password,
		birthday,
		create_at,
		update_at,
		deleted_at
	FROM
		users
	WHERE 
		id = $1 `

	row := user.DB.QueryRow(query, id) //sorovni bajarish uchun queryrow dan foydalanib qatorni qaytaradi
	err := row.Scan(                   //qaytadigan malumotlarni scan qilib oqib oladi va models.Userni users ozgaruvchiga saqlab qoyadi.
		&users.UserID,
		&users.Name,
		&users.Email,
		&users.Birthday,
		&users.Password,
		&users.CreatedAt,
		&users.UpdatedAt,
		&users.DeletedAt,
	)

	if err != nil {
		log.Println("error get user by id: ", err)
	}

	return &users, nil
}

func (db *UserRepo) GetAllUsers(name string) (*[]models.User, error) {
	resp := []models.User{}

	query := `SELECT * FROM users WHERE deleted_at = 0`
	var args []interface{}

	if name != "" {
		query += " AND name = ?"
		args = append(args, name)
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		resp = append(resp, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (db *UserRepo) Search(email string) (*[]models.User, error) {
	var users []models.User

	baseQuery := `
        SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at 
        FROM users 
        WHERE deleted_at = 0`

	var conditions []string
	var args []interface{}

	if email != "" {
		conditions = append(conditions, "email = ?")
		args = append(args, email)
	}

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	rows, err := db.DB.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}

func (db *UserRepo) UpdatePassword(id string, newPassword string) error {

	_, err := db.DB.Exec(
		`update users set password=$1 where user_id=$2 and deleted_at=0`, newPassword, id,
	)
	return err
}

func (db *UserRepo) DeleteUserById(id string) error {

	_, err := db.DB.Exec(
		`update users set deleted_at = date_part('epoch', current_timestamp)::INT where user_id=$1`, id,
	)
	return err
}

//date_part('epoch', current_timestamp)::INT -> 13245678908765
