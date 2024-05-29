
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Product struct {
	ID    int
	Name  string
	Price int
}

type Order struct {
	ID        int
	OrderDate string
}

type OrderProduct struct {
	ID         int
	OrderID    int
	ProductID  int
	Quantity   int
}

func main() {
	connStr := "user=postgres password=root dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// addData(db)

	fetchData(db)

	displayData(db)
}

func addData(db *sql.DB) {
	_, err := db.Exec(`
		INSERT INTO Product (name, price) VALUES
		('Product 1', 100),
		('Product 2', 150),
		('Product 3', 200)
	`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		INSERT INTO orders (order_date) VALUES
		('2024-01-12'),
		('2024-01-13'),
		('2024-01-14')
	`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		INSERT INTO order_product (order_id, product_id, quantity) VALUES
		(1, 1, 3),
		(1, 2, 2),
		(1, 3, 1)
	`)
	if err != nil {
		panic(err)
	}
}

func fetchData(db *sql.DB) {
    rows, err := db.Query("SELECT id, name, price FROM Product")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    fmt.Println("Products:")
    for rows.Next() {
        var product Product
        if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("ID: %d, Name: %s, Price: %d\n", product.ID, product.Name, product.Price)
    }

    rows, err = db.Query("SELECT id, order_date FROM orders")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    fmt.Println("\nOrders:")
    for rows.Next() {
        var order Order
        if err := rows.Scan(&order.ID, &order.OrderDate); err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("ID: %d, OrderDate: %s\n", order.ID, order.OrderDate)
    }

    rows, err = db.Query("SELECT id, order_id, product_id, quantity FROM order_product")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    fmt.Println("\nOrder_Product:")
    for rows.Next() {
        var orderProduct OrderProduct
        if err := rows.Scan(&orderProduct.ID, &orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Quantity); err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("ID: %d, OrderID: %d, ProductID: %d, Quantity: %d\n", orderProduct.ID, orderProduct.OrderID, orderProduct.ProductID, orderProduct.Quantity)
    }
}

func displayData(db *sql.DB) {
    fetchData(db)
}
