create table users (
    user_id uuid primary key default gen_random_uuid(),
    name varchar(100),
    email varchar(100),
    birthday timestamp,
    password varchar(100),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at bigint default 0
);

create table courses (
    course_id uuid primary key default gen_random_uuid() not null,
    title varchar(100),
    description varchar(100),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at bigint default 0
);

create table lessons (
    lesson_id uuid primary key default gen_random_uuid(),
    course_id uuid references courses(course_id),
    title varchar(100),
    content text,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at bigint default 0
);

create table enrollments (
    enrollment_id uuid primary key default gen_random_uuid(),
    user_id uuid references users(user_id),
    course_id uuid references courses(course_id),
    enrollment_date timestamp,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at bigint default 0
);