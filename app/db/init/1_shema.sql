CREATE TABLE users
(
    user_id           SERIAL,
    user_name         VARCHAR(255) NOT NULL,
    user_display_name VARCHAR(255) NOT NULL
);

CREATE TABLE books
(
    book_id    SERIAL,
    user_id    BIGINT UNSIGNED NOT NULL,
    isbn13     INTEGER         NOT NULL,
    book_title VARCHAR(255)    NULL,
    PRIMARY KEY (book_id),
    FOREIGN KEY (user_id) REFERENCEs users(user_id)
);


