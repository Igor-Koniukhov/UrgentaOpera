package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pasztorpisti/qs"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"todo/backend/internal/configs"
	"todo/backend/internal/models"
	"todo/backend/internal/repository"
)

type TicketI interface {
	CreateTicket(w http.ResponseWriter, r *http.Request)
	GetTicketById(w http.ResponseWriter, r *http.Request)
	GetAllTickets(w http.ResponseWriter, r *http.Request)
	UpdateTicket(w http.ResponseWriter, r *http.Request)
	DeleteTicket(w http.ResponseWriter, r *http.Request)
}
type Ticket struct {
	App  *configs.AppConfig
	repo repository.Repository
}
type GetTicketsByUserID struct {
	Q1 string `qs:"user_id"`
	Q2 string `qs:"ticket"`
}

func NewTicket(app *configs.AppConfig, repo repository.Repository) *Ticket {
	return &Ticket{
		App:  app,
		repo: repo,
	}
}

func (t Ticket) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket *models.Ticket
	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ticket)
	_, id, err := t.repo.CreateTicket(ticket)
	if err != nil {
		fmt.Println(err)
	}
	message := make(map[string]interface{})
	message["ticketId"] = id
	err = json.NewEncoder(w).Encode(&message)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}
func NewGetTicketsByUserID(query url.Values) (req *GetTicketsByUserID, err error) {
	q, err := qs.Marshal(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = qs.Unmarshal(&req, q)
	if err != nil {
		return nil, err
	}

	if req.Q1 == "" || req.Q2 == "" {
		req.Q1 = "1"
		req.Q2 = "1"
		return req, errors.New(" Invalid url string parameters!")
	}

	return &GetTicketsByUserID{
		Q1: req.Q1,
		Q2: req.Q2,
	}, err
}
func (t Ticket) GetAllTickets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		return
	}

	tickets, err := t.repo.GetAllTickets(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewEncoder(w).Encode(&tickets)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (t Ticket) GetTicketById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		return
	}
	ticket, err := t.repo.GetTicketById(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewEncoder(w).Encode(&ticket)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

func (t Ticket) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		fmt.Println(err)
	}
	updatedTicket, err := t.repo.UpdateTicket(ticket)
	if err != nil {
		fmt.Println(err)
	}
	err = json.NewEncoder(w).Encode(&updatedTicket)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

func (t Ticket) DeleteTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.repo.DeleteTicket(id)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusAccepted)
}
