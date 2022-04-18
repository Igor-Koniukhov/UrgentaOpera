-- +goose Up
CREATE TABLE IF NOT EXISTS tickets
(
    id         INT AUTO_INCREMENT,
    list       INT,
    user_id    INT,
    title      VARCHAR(100),
    background VARCHAR(45),
    color      VARCHAR(45),
    status     VARCHAR(45),
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT tickets_pk
        PRIMARY KEY (id),
    CONSTRAINT tickets_users_id_fk
        FOREIGN KEY (user_id) REFERENCES users (id)
            ON DELETE CASCADE
);
-- +goose Down
DROP TABLE tickets;
