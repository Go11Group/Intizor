create table userlar(
    id    uuid primary key not null default gen_random_uuid(),
    name  varchar   not null,
    age int not null
);

INSERT INTO userlar (name, age) VALUES
('Alice', 30),
('Bob', 24),
('Charlie', 35),
('Diana', 28),
('Eve', 22),
('Frank', 45),
('Grace', 33),
('Hank', 29),
('Ivy', 31),
('Jack', 26);


create table car
(
    id     uuid primary key not null default gen_random_uuid(),
    name   varchar          not null,
    year    int not null,
    color varchar not null 
);

INSERT INTO car (name, year, color) VALUES
('Toyota Corolla', 2020, 'Blue'),
('Honda Civic', 2018, 'Red'),
('Ford Mustang', 2021, 'Black'),
('Chevrolet Impala', 2017, 'White'),
('BMW 3 Series', 2019, 'Silver'),
('Audi A4', 2020, 'Gray'),
('Tesla Model 3', 2021, 'Black'),
('Nissan Altima', 2018, 'Blue'),
('Hyundai Elantra', 2019, 'White'),
('Kia Optima', 2021, 'Red');


create table users_cars(
    id serial primary key,
    car_id uuid not null,
    user_id uuid not null,
    foreign key(car_id) references car(id),
    foreign key(user_id) references user(id)
);