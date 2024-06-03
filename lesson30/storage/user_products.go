package storage

import "database/sql"

type UserProductRepo struct {
    db *sql.DB
}

func NewUserProductRepo(db *sql.DB) *UserProductRepo {
    return &UserProductRepo{db}
}

func (r *UserProductRepo) Create(userID, productID int) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec("INSERT INTO user_products (user_id, product_id) VALUES ($1, $2)", userID, productID)
    if err != nil {
        return err
    }

    return tx.Commit()
}