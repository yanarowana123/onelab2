CREATE TABLE IF NOT EXISTS books
(
    id         uuid primary key,
    name       varchar(255) not null,
    author_id     uuid         not null,
    created_at timestamp default now(),
    CONSTRAINT fk_user_id
        FOREIGN KEY (author_id)
            REFERENCES users (id)
            ON UPDATE CASCADE ON DELETE CASCADE
)


