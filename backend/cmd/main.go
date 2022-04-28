package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"todo/driver"
	"todo/internal/configs"
	"todo/internal/handlers"
	"todo/internal/repository"
	"todo/internal/server"
)

func main() {
	app := configs.NewAppConfig()
	Dr, err := driver.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
	defer Dr.SQL.Close()
	rep := repository.NewRepository(app, Dr.SQL)
	www := handlers.NewHandlerStruct(app, *rep)
	srv := &server.Server{}
	go func() {
		err := srv.Serve(
			os.Getenv("PORT"),
			routes(www),
		)
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	fmt.Println(" TODO APPLICATION STARTED ON PORT" +
		os.Getenv("PORT") +
		os.Getenv("CONSOLE_NUV_API"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("TODO APPLICATION SHUTTING DOWN")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

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
