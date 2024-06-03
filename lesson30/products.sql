-- Products jadvali
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock_quantity INT NOT NULL
);

INSERT INTO products (name, description, price, stock_quantity) VALUES
('Product 1', 'Description of Product 1', 19.99, 100),
('Product 2', 'Description of Product 2', 29.99, 50),
('Product 3', 'Description of Product 3', 9.99, 200),
('Product 4', 'Description of Product 4', 49.99, 75),
('Product 5', 'Description of Product 5', 14.99, 150),
('Product 6', 'Description of Product 6', 39.99, 80),
('Product 7', 'Description of Product 7', 24.99, 120),
('Product 8', 'Description of Product 8', 34.99, 90),
('Product 9', 'Description of Product 9', 64.99, 30),
('Product 10', 'Description of Product 10', 19.99, 100);


select*from products;

 id |    name    |        description        | price | stock_quantity 
----+------------+---------------------------+-------+----------------
  1 | Product 1  | Description of Product 1  | 19.99 |            100
  2 | Product 2  | Description of Product 2  | 29.99 |             50
  3 | Product 3  | Description of Product 3  |  9.99 |            200
  4 | Product 4  | Description of Product 4  | 49.99 |             75
  5 | Product 5  | Description of Product 5  | 14.99 |            150
  6 | Product 6  | Description of Product 6  | 39.99 |             80
  7 | Product 7  | Description of Product 7  | 24.99 |            120
  8 | Product 8  | Description of Product 8  | 34.99 |             90
  9 | Product 9  | Description of Product 9  | 64.99 |             30
 10 | Product 10 | Description of Product 10 | 19.99 |            100
(10 rows)
