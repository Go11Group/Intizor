
-- inner join

-- table yaratish
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    author_id INT
);

-- table yaratish
CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- tablega ma'lumot yozish
INSERT INTO authors (name) VALUES 
('Erkin Vohidov'),
('Togay Murod');
('Otkir Hoshimov'),
('Tohir Malik'),
('Navoiy'),
('Bobur'),
('Abdulla Qodiriy')

-- tablega ma'lumot yozish
INSERT INTO books (name, author_id) VALUES
('Ikki eshik orasi', 3),
('Boburnoma', 6),
('Hamsa', 5),
('Ot kishnagan oqshom', 2);
('Roman', 3),
('Dunyoning ishlari', 3),
('Oydinda yurgan odamlar', 3),
('Shaytanat', 4);
('Nido', 1),
('Otkan kunlar', 7),
('Mehrobdan chayon', 7);

SELECT b.id, b.name, a.name FROM books b INNER JOIN authors a ON b.author_id = a.id ORDER BY b.name DESC;
SELECT b.id, b.name, a.name FROM authors a INNER JOIN books b ON b.author_id = a.id WHERE b.id = 2;

================================================================

--left join 

CREATE TABLE IF NOT EXISTS film (
    film_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL
);

-- inventory jadvali
CREATE TABLE IF NOT EXISTS inventory (
    inventory_id SERIAL PRIMARY KEY,
    film_id INT,
    -- Qo'shimcha ustunlar uchun kerakli ma'lumotlarni qo'shing
    FOREIGN KEY (film_id) REFERENCES film(film_id)
);

-- film jadviga ma'lumot qo'shish
INSERT INTO film (title)
VALUES
    ('Film1'),
    ('Film2'),
    ('Film3');

-- inventory jadviga ma'lumot qo'shish
INSERT INTO inventory (film_id)
VALUES
    (1),
    (2);


SELECT
	film.film_id,
	title,
	inventory_id
FROM
	film
LEFT JOIN inventory 
    ON inventory.film_id = film.film_id
ORDER BY title;
-- 
SELECT
	film.film_id,
	film.title,
	inventory_id
FROM
	film
LEFT JOIN inventory 
   ON inventory.film_id = film.film_id
WHERE inventory.film_id IS NULL
ORDER BY title;
==================================================================

-- right join

-- table yaratish
CREATE TABLE films(
   film_id SERIAL PRIMARY KEY,
   title varchar(255) NOT NULL
);

-- tablega ma'lumot yozish
INSERT INTO films(title)
VALUES('Joker'),
      ('Avengers: Endgame'),
      ('Parasite');

-- yable yaratish
CREATE TABLE film_reviews(
   review_id SERIAL PRIMARY KEY,
   film_id INT,
   review VARCHAR(255) NOT NULL	
);

-- tablega ma'lumot yozish
INSERT INTO film_reviews(film_id, review)
VALUES(1, 'Excellent'),
      (1, 'Awesome'),
      (2, 'Cool'),
      (NULL, 'Beautiful');


SELECT 
   review, 
   title
FROM 
   films
RIGHT JOIN film_reviews 
   ON film_reviews.film_id = films.film_id;
-- 
SELECT review, title
FROM films
RIGHT JOIN film_reviews USING (film_id);
-- 
SELECT review, title
FROM films
RIGHT JOIN film_reviews USING (film_id)
WHERE title IS NULL;

================================================================================================================================

-- full join

-- table yaratish

CREATE TABLE departments (
	department_id serial PRIMARY KEY,
	department_name VARCHAR (255) NOT NULL
);

CREATE TABLE employees (
	employee_id serial PRIMARY KEY,
	employee_name VARCHAR (255),
	department_id INTEGER
);

-- tablega ma'lumot yozish
INSERT INTO departments (department_name)
VALUES
	('Sales'),
	('Marketing'),
	('HR'),
	('IT'),
	('Production');

INSERT INTO employees (
	employee_name,
	department_id
)
VALUES
	('Bette Nicholson', 1),
	('Christian Gable', 1),
	('Joe Swank', 2),
	('Fred Costner', 3),
	('Sandra Kilmer', 4),
	('Julia Mcqueen', NULL);


SELECT
	employee_name,
	department_name
FROM
	employees e
FULL OUTER JOIN departments d 
        ON d.department_id = e.department_id;
-- 
SELECT
	employee_name,
	department_name
FROM
	employees e
FULL OUTER JOIN departments d 
        ON d.department_id = e.department_id
WHERE
	employee_name IS NULL;




    -- leetcode 584

-- table yaratish
CREATE TABLE customers(
    id int,
    name VARCHAR,
    referee_id int
);

-- tablega ma'lumot yozish
INSERT INTO customers(id, name, referee_id) 
VALUES (1, 'Will',null),
(2, 'Jeyn',null),
(3, 'Alex',2),
(4, 'Bill',null),
(5, 'Zak',1),
(6, 'Mark',2);

-- tabledan ma'lumotlarni olish
SELECT * FROM customers;
SELECT name FROM customers
WHERE referee_id != 2 OR referee_id IS NULL;

================================================================
--leetcode 1148

-- table yaratish
CREATE TABLE Views(
    article_id int,
    author_id int,
    viewer_id int,
    view_date date
);

-- tablega ma'lumot yozish
INSERT INTO Views(article_id, author_id, viewer_id, view_date) 
VALUES (1, 3, 5, '2019-08-01'),
(1, 3, 6, '2019-08-02'),
(2, 7, 7, '2019-08-01'),
(2, 7, 6, '2019-08-02'),
(4, 7, 1, '2019-07-22'),
(3, 4, 4, '2019-07-21'),
(3, 4, 4, '2019-07-21');

-- tabledan ma'lumotlarni olish
SELECT * FROM Views;
SELECT DISTINCT author_id AS id
FROM Views
WHERE author_id = viewer_id
ORDER BY id ASC;

==============================================================

--leetcode 1683

-- table yaratish
CREATE TABLE Tweets(
    tweet_id int,
    content VARCHAR
);

-- tablega ma'lumot yozish
INSERT INTO Tweets(tweet_id, content) 
VALUES (1, 'Vote for Biden'),
(2, 'Let us make Amerika great again!');

-- tabledan ma'lumotlarni olish
SELECT * FROM Tweets;
SELECT tweet_id 
FROM Tweets
WHERE LENGTH(content) > 15;
