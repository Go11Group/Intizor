CREATE TABLE IF NOT EXISTS users1 (
                id SERIAL PRIMARY KEY,
                first_name VARCHAR(50),
                last_name VARCHAR(50),
                email VARCHAR(100),
                gender VARCHAR(10),
                age INT
        );


INSERT INTO users1 (first_name, last_name, email, gender, age) VALUES
                                                    
('John', 'Doe', 'john.doe@example.com', 'Male', 28),
('Jane', 'Smith', 'jane.smith@example.com', 'Female', 32),
('Alice', 'Johnson', 'alice.johnson@example.com', 'Female', 24),
('Bob', 'Brown', 'bob.brown@example.com', 'Male', 45),
('Charlie', 'Davis', 'charlie.davis@example.com', 'Male', 36),
('Dana', 'Evans', 'dana.evans@example.com', 'Female', 29),
('Eve', 'Franklin', 'eve.franklin@example.com', 'Female', 41),
('Frank', 'Green', 'frank.green@example.com', 'Male', 37),
('Grace', 'Hill', 'grace.hill@example.com', 'Female', 26),
('Hank', 'Ingram', 'hank.ingram@example.com', 'Male', 33);



select*from users1;
 id | first_name | last_name |           email           | gender | age 
----+------------+-----------+---------------------------+--------+-----
  1 | John       | Doe       | john.doe@example.com      | Male   |  28
  2 | Jane       | Smith     | jane.smith@example.com    | Female |  32
  3 | Alice      | Johnson   | alice.johnson@example.com | Female |  24
  4 | Bob        | Brown     | bob.brown@example.com     | Male   |  45
  5 | Charlie    | Davis     | charlie.davis@example.com | Male   |  36
  6 | Dana       | Evans     | dana.evans@example.com    | Female |  29
  7 | Eve        | Franklin  | eve.franklin@example.com  | Female |  41
  8 | Frank      | Green     | frank.green@example.com   | Male   |  37
  9 | Grace      | Hill      | grace.hill@example.com    | Female |  26
 10 | Hank       | Ingram    | hank.ingram@example.com   | Male   |  33
