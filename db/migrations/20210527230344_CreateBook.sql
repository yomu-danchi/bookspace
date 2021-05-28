-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE books (
    book_id SERIAL,
    owner_id INTEGER NOT NULL,
    isbn13 INTEGER NOT NULL,
    book_title VARCHAR(255) NULL,
    primary key (book_id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE books;