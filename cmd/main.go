package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
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
