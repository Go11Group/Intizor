package storage

import (
	"database/sql"
	"github.com/Go11Group/Intizor/lesson30/model"
)


type ProductRepo struct {
    db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
    return &ProductRepo{db}
}

func (r *ProductRepo) Create(product *model.Product) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec("INSERT INTO products (name, description, price, stock_quantity) VALUES ($1, $2, $3, $4)", product.Name, product.Description, product.Price, product.StockQuantity)
    if err != nil {
        return err
    }

    return tx.Commit()
}

func (r *ProductRepo) Get(id int) (*model.Product, error) {
    var product model.Product
    err := r.db.QueryRow("SELECT id, name, description, price, stock_quantity FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.StockQuantity)
    if err != nil {
        return nil, err
    }
    return &product, nil
}

func (r *ProductRepo) Update(product *model.Product) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec("UPDATE products SET name = $1, description = $2, price = $3, stock_quantity = $4 WHERE id = $5", product.Name, product.Description, product.Price, product.StockQuantity, product.ID)
    if err != nil {
        return err
    }

    return tx.Commit()
}

func (r *ProductRepo) Delete(id int) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec("DELETE FROM products WHERE id = $1", id)
    if err != nil {
        return err
    }

    return tx.Commit()
}