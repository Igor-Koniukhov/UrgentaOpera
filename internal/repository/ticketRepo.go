package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"todo/internal/configs"
	"todo/internal/models"
)

type TicketRepository interface {
	CreateTicket(t *models.Ticket) (*models.Ticket, int, error)
	GetAllTickets() ([]models.Ticket, error)
	GetTicketById(id int) (models.Ticket, error)
	UpdateTicket(t models.Ticket) (models.Ticket, error)
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

func (tr TicketRepo) CreateTicket(t *models.Ticket) (*models.Ticket, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO %s (list, title, background, color, status)"+
		" VALUES(?,?,?,?,?) ", models.TicketsTable)
	res, err := tr.DB.ExecContext(ctx, sqlStmt,
		t.List,
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
func (tr TicketRepo) GetAllTickets() ([]models.Ticket, error) {
	var ticket models.Ticket
	var tickets []models.Ticket
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT id, list, title, background, color,"+
		"status, deleted_at, created_at, updated_at FROM %s ",
		models.TicketsTable)
	results, err := tr.DB.QueryContext(ctx, sqlStmt)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for results.Next() {
		err = results.Scan(
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
			return nil, err
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

func (tr TicketRepo) GetTicketById(id int) (models.Ticket, error) {
	var ticket models.Ticket
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT id, list, title, background, color,"+
		"status, deleted_at, created_at, updated_at FROM %s WHERE id = ? ",
		models.TicketsTable)
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
		return models.Ticket{}, err
	}
	return ticket, nil
}

func (tr TicketRepo) UpdateTicket(t models.Ticket) (models.Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET  list=?, title=?, "+
		"background=? ,color=?, status=?  WHERE id=? ", models.TasksTable)
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
		return models.Ticket{}, nil
	}
	return t, nil
}

func (tr TicketRepo) DeleteTicket(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("UPDATE %s SET deleted_at=?, status=? WHERE id=?", models.TicketsTable)
	_, err := tr.DB.ExecContext(ctx, sqlStmt, time.Now(), "deleted", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
