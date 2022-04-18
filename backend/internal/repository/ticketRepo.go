package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"todo/backend/internal/configs"
	models2 "todo/backend/internal/models"
)

type TicketRepository interface {
	CreateTicket(t *models2.Ticket) (*models2.Ticket, int, error)
	GetAllTickets(userId int) ([]models2.Ticket, error)
	GetTicketById(id int) (models2.Ticket, error)
	UpdateTicket(t models2.Ticket) (models2.Ticket, error)
	DeleteTicket(id int) error
}

type TicketRepo struct {
	App *configs.AppConfig
	DB  *sql.DB
}

func NewTicketRepo(app *configs.AppConfig, db *sql.DB) *TicketRepo {
	return &TicketRepo{
		App: app,
		DB:  db,
	}
}

func (tr TicketRepo) CreateTicket(t *models2.Ticket) (*models2.Ticket, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO %s (list, user_id, title, background, color, status)"+
		" VALUES(?, ?, ?,?,?,?) ", models2.TicketsTable)
	res, err := tr.DB.ExecContext(ctx, sqlStmt,
		t.List,
		t.UserId,
		t.Title,
		t.Background,
		t.Color,
		t.Status,
	)

	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}
	ticketId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}
	return t, int(ticketId), nil
}
func (tr TicketRepo) GetAllTickets(userId int) ([]models2.Ticket, error) {
	var ticket models2.Ticket
	var tickets []models2.Ticket
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT id, list, user_id, title, background, color,"+
		"status, deleted_at, created_at, updated_at FROM %s WHERE user_id=?",
		models2.TicketsTable)
	results, err := tr.DB.QueryContext(ctx, sqlStmt, userId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for results.Next() {
		err = results.Scan(
			&ticket.ID,
			&ticket.List,
			&ticket.UserId,
			&ticket.Title,
			&ticket.Background,
			&ticket.Color,
			&ticket.Status,
			&ticket.DeletedAt,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

func (tr TicketRepo) GetTicketById(id int) (models2.Ticket, error) {
	var ticket models2.Ticket
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT id, list, title, background, color,"+
		"status, deleted_at, created_at, updated_at FROM %s WHERE id = ? ",
		models2.TicketsTable)
	row := tr.DB.QueryRowContext(ctx, sqlStmt, id)
	err := row.Scan(
		&ticket.ID,
		&ticket.List,
		&ticket.Title,
		&ticket.Background,
		&ticket.Color,
		&ticket.Status,
		&ticket.DeletedAt,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err)
		return models2.Ticket{}, err
	}
	return ticket, nil
}

func (tr TicketRepo) UpdateTicket(t models2.Ticket) (models2.Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET  list=?, title=?, "+
		"background=? ,color=?, status=?  WHERE id=? ", models2.TasksTable)
	_, err := tr.DB.ExecContext(ctx, sqlStmt,
		&t.List,
		&t.Title,
		&t.Background,
		&t.Color,
		&t.Status,
		&t.ID,
	)
	if err != nil {
		fmt.Println(err)
		return models2.Ticket{}, nil
	}
	return t, nil
}

func (tr TicketRepo) DeleteTicket(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET deleted_at=?, status=? WHERE id=?", models2.TicketsTable)
	_, err := tr.DB.ExecContext(ctx, sqlStmt, time.Now(), "deleted", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
