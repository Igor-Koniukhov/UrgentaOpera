package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"todo/internal/handlers"
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

	return r
}
