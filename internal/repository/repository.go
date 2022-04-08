package repository

import (
	"database/sql"
	"todo/internal/configs"
)

type Repository struct {
	UserRepository
}

func NewRepository(app *configs.AppConfig, db *sql.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepo(app, db),
	}
}
