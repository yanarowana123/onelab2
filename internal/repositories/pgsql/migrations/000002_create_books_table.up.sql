CREATE TABLE IF NOT EXISTS books
(
    id       uuid primary key,
    name    varchar(255) not null,
    author varchar(255) not null,
    created_at timestamp default now()
)


