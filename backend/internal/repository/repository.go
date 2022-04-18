package repository

import (
	"database/sql"
	"todo/backend/internal/configs"
)

type Repository struct {
	UserRepository
	TicketRepository
	TaskRepository
}

func NewRepository(app *configs.AppConfig, db *sql.DB) *Repository {
	return &Repository{
		UserRepository:   NewUserRepo(app, db),
		TicketRepository: NewTicketRepo(app, db),
		TaskRepository:   NewTaskRepo(app, db),
	}
}
