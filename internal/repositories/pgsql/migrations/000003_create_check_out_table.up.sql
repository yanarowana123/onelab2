CREATE TABLE IF NOT EXISTS check_out
(
    user_id uuid not null,
    book_id uuid not null,
    checked_out_at timestamp default now(),
    returned_at timestamp default null,
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_book_id
        FOREIGN KEY(book_id)
        REFERENCES books(id)
        ON UPDATE CASCADE ON DELETE CASCADE
)
