package storage

import (
	"database/sql"
	"fmt"

	"github.com/Go11Group/Intizor/lesson30/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) Create(user *model.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	_, err = tx.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
    
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed %v", err)
	}
	return nil
}

func (r *UserRepo) Get(id int) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id, username, email, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) Update(user *model.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4", user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *UserRepo) Delete(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
