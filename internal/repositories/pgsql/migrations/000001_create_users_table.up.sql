CREATE TABLE IF NOT EXISTS users
(
    id       uuid primary key,
    login    varchar(255) not null unique,
    name     varchar(255) not null,
    password varchar(255) not null,
    created_at timestamp default now()
)


