-- Users jadvali
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);

INSERT INTO users (username, email, password) VALUES
('user1', 'user1@example.com', 'password1'),
('user2', 'user2@example.com', 'password2'),
('user3', 'user3@example.com', 'password3'),
('user4', 'user4@example.com', 'password4'),
('user5', 'user5@example.com', 'password5'),
('user6', 'user6@example.com', 'password6'),
('user7', 'user7@example.com', 'password7'),
('user8', 'user8@example.com', 'password8'),
('user9', 'user9@example.com', 'password9'),
('user10', 'user10@example.com', 'password10');



select* from users;

id | username |       email        |  password  
----+----------+--------------------+------------
  1 | user1    | user1@example.com  | password1
  2 | user2    | user2@example.com  | password2
  3 | user3    | user3@example.com  | password3
  4 | user4    | user4@example.com  | password4
  5 | user5    | user5@example.com  | password5
  6 | user6    | user6@example.com  | password6
  7 | user7    | user7@example.com  | password7
  8 | user8    | user8@example.com  | password8
  9 | user9    | user9@example.com  | password9
 10 | user10   | user10@example.com | password10
(10 rows)
