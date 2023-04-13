CREATE TABLE IF NOT EXISTS users
(
    id       uuid primary key,
    email    varchar(255) not null unique,
    first_name     varchar(255) not null,
    last_name     varchar(255) not null,
    password varchar(255) not null,
    created_at timestamp default now()
)


