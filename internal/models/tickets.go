package models

import "time"

const TicketsTable = "tickets"

type Ticket struct {
	ID         int       `json:"id"`
	List       int       `json:"list"`
	Title      string    `json:"title"`
	Background string    `json:"background"`
	Color      string    `json:"color"`
	Status     string    `json:"status"`
	DeletedAt  time.Time `json:"deleted_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
