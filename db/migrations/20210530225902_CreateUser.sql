
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Users
(
    user_id           SERIAL,
    user_name         VARCHAR(255) NOT NULL,
    user_display_name VARCHAR(255) NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
