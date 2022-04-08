-- +goose Up
CREATE TABLE IF NOT EXISTS users
(
    id         INT AUTO_INCREMENT,
    name       VARCHAR(45),
    phone      VARCHAR(90),
    email      VARCHAR(45),
    password   VARCHAR(90),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT users_pk
        PRIMARY KEY (id)
);
-- +goose Down
DROP TABLE users CASCADE;
