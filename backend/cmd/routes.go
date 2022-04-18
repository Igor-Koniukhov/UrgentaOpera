package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"todo/backend/internal/handlers"
)

func routes(www *handlers.HandlerStruct) http.Handler {
	r := mux.NewRouter()
	r.Use(www.JSON, www.CORS)
	r.HandleFunc("/login", www.UserI.Login).Methods("POST")
	r.HandleFunc("/logout", www.UserI.LogOut).Methods("GET")
	r.HandleFunc("/registration", www.UserI.Registration).Methods("POST")
	r.HandleFunc("/forgot", www.UserI.Forgot).Methods("POST")
	r.HandleFunc("/reset", www.UserI.Reset).Methods("POST")
	r.HandleFunc("/users", www.UserI.GetAllUsers).Methods("GET")

	r.HandleFunc("/create-ticket", www.TicketI.CreateTicket).Methods("POST")
	r.HandleFunc("/ticket/{id:[0-9]}", www.TicketI.GetTicketById).Methods("GET")
	r.HandleFunc("/tickets/{id:[0-9]}", www.TicketI.GetAllTickets).Methods("GET")
	r.HandleFunc("/update-ticket/{id:[0-9]}", www.TicketI.UpdateTicket).Methods("PUT")
	r.HandleFunc("/delete-ticket/{id:[0-9]}", www.TicketI.DeleteTicket).Methods("DELETE")

	r.HandleFunc("/create-task", www.TaskI.CreateTask).Methods("POST")
	r.HandleFunc("/task/{id:[0-9]}", www.TaskI.GetTaskById).Methods("GET")
	r.HandleFunc("/task-status/{id:[0-9]}", www.TaskI.SetTaskStatus).Methods("PUT")
	r.HandleFunc("/update-task/{id:[0-9]}", www.TaskI.UpdateTask).Methods("PUT")
	r.HandleFunc("/delete-task/{id:[0-9]}", www.TaskI.DeleteTask).Methods("DELETE")

	return r
}
