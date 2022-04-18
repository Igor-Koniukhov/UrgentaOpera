package models

import "time"

const TasksTable = "tasks"

type Task struct {
	ID        int       `json:"id"`
	TicketId  int       `json:"ticket_id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	DueDone   bool      `json:"dueDone"`
	Status    string    `json:"status"`
	DeletedAt time.Time `json:"deleted_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
