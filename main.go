package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pushpendras21/go-server/config"
	"github.com/pushpendras21/go-server/db"
	"github.com/pushpendras21/go-server/routes"
)

func main() {

	config.Config.LoadConfig()
	handler := routes.MounteRoutes()

	fmt.Printf("Starting server on %s", config.Config.AppPort)

	db.InitDB()
	server := &http.Server{
		Addr:    config.Config.AppPort,
		Handler: handler,
	}
	defer db.DB.Close(context.Background())
	server.ListenAndServe()

}
