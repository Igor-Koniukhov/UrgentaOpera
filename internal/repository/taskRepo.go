package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"todo/internal/configs"
	"todo/internal/models"
)

type TaskRepository interface {
	CreateTask(t *models.Task) (*models.Task, int, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id int) (models.Task, error)
	UpdateTask(t models.Task) (models.Task, error)
	SetTaskStatus(id int, status string) error
	DeleteTask(id int) error
}
type TaskRepo struct {
	App   *configs.AppConfig
	DB    *sql.DB
	task  models.Task
	tasks []models.Task
}

func NewTaskRepo(app *configs.AppConfig, db *sql.DB) *TaskRepo {
	return &TaskRepo{
		App: app,
		DB:  db,
	}
}

func (tr TaskRepo) CreateTask(t *models.Task) (*models.Task, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO %s (ticket_id, title, done, DueDone)"+
		" VALUES(?,?,?,?) ", models.TasksTable)

	res, err := tr.DB.ExecContext(ctx, sqlStmt,
		t.TicketId,
		t.Title,
		t.Done,
		t.DueDone,
	)
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}
	taskId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}
	return t, int(taskId), nil
}

func (tr TaskRepo) GetTaskById(id int) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT id, ticket_id, title, done, dueDone,"+
		", status, deleted_at, created_at, updated_at FROM %s WHERE id = ? ",
		models.TasksTable)
	row := tr.DB.QueryRowContext(ctx, sqlStmt, id)
	err := row.Scan(
		&tr.task.ID,
		&tr.task.TicketId,
		&tr.task.Title,
		&tr.task.Done,
		&tr.task.DueDone,
		&tr.task.Status,
		&tr.task.DeletedAt,
		&tr.task.CreatedAt,
		&tr.task.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err)
		return models.Task{}, err
	}
	return tr.task, nil
}

func (tr TaskRepo) UpdateTask(t models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET  ticket_id=?, title=?, "+
		"Done=? ,dueDone=?, status=? , updated_at, WHERE id=? ", models.TasksTable)
	_, err := tr.DB.ExecContext(ctx, sqlStmt,
		t.TicketId,
		t.Title,
		t.Done,
		t.DueDone,
		"updated",
		time.Now(),
		t.ID,
	)
	if err != nil {
		fmt.Println(err)
		return models.Task{}, nil
	}
	return t, nil
}

func (tr TaskRepo) DeleteTask(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET deleted_at=?, status=? WHERE id=?", models.TasksTable)
	_, err := tr.DB.ExecContext(ctx, sqlStmt, time.Now(), "deleted", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (tr TaskRepo) SetTaskStatus(id int, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET done=?, status=? WHERE id=?", models.TasksTable)
	_, err := tr.DB.ExecContext(ctx, sqlStmt, true, status, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
