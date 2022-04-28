package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"todo/internal/configs"
	models2 "todo/internal/models"
)

type TaskRepository interface {
	CreateTask(t *models2.Task) (*models2.Task, int, error)
	GetAllTasks() ([]models2.Task, error)
	GetTaskById(id int) (models2.Task, error)
	UpdateTask(t models2.Task) (models2.Task, error)
	SetTaskStatus(id int, status string) error
	DeleteTask(id int) error
}
type TaskRepo struct {
	App   *configs.AppConfig
	DB    *sql.DB
	task  models2.Task
	tasks []models2.Task
}

func NewTaskRepo(app *configs.AppConfig, db *sql.DB) *TaskRepo {
	return &TaskRepo{
		App: app,
		DB:  db,
	}
}

func (tr TaskRepo) CreateTask(t *models2.Task) (*models2.Task, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO %s (ticket_id, title, done, DueDone)"+
		" VALUES(?,?,?,?) ", models2.TasksTable)

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
func (tr TaskRepo) GetAllTasks() ([]models2.Task, error) {
	var task models2.Task
	var tasks []models2.Task
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT id, ticket_id, title, done, dueDone,"+
		"status, deleted_at, created_at, updated_at FROM %s ",
		models2.TicketsTable)
	results, err := tr.DB.QueryContext(ctx, sqlStmt)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for results.Next() {
		err = results.Scan(
			&task.ID,
			&task.TicketId,
			&task.Title,
			&task.Done,
			&task.DueDone,
			&task.Status,
			&task.DeletedAt,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (tr TaskRepo) GetTaskById(id int) (models2.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT id, ticket_id, title, done, dueDone,"+
		", status, deleted_at, created_at, updated_at FROM %s WHERE id = ? ",
		models2.TasksTable)
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
		return models2.Task{}, err
	}
	return tr.task, nil
}

func (tr TaskRepo) UpdateTask(t models2.Task) (models2.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET  ticket_id=?, title=?, "+
		"Done=? ,dueDone=?, status=? , updated_at, WHERE id=? ", models2.TasksTable)
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
		return models2.Task{}, nil
	}
	return t, nil
}

func (tr TaskRepo) DeleteTask(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET deleted_at=?, status=? WHERE id=?", models2.TasksTable)
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
	sqlStmt := fmt.Sprintf("UPDATE %s SET done=?, status=? WHERE id=?", models2.TasksTable)
	_, err := tr.DB.ExecContext(ctx, sqlStmt, true, status, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
