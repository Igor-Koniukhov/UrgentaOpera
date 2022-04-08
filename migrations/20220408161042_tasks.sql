-- +goose Up
CREATE TABLE IF NOT EXISTS tasks
(
    id         INT AUTO_INCREMENT,
    ticket_id  INT,
    title      LONGTEXT,
    dueDone    TINYINT,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT tasks_pk
        PRIMARY KEY (id),
    CONSTRAINT tasks_tickets_id_fk
        FOREIGN KEY (ticket_id) REFERENCES tickets (id)
            ON DELETE CASCADE
);
-- +goose Down
DROP TABLE tasks;
