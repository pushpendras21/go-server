package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn
var err error

func InitDB() {
	//urlExample := "host=localhost port=5432 user=postgres password=password dbname=concurrency sslmode=disable timezone=UTC connect_timeout=5"
	urlExample := "postgres://postgres:adminPassword@localhost:5432/tasks"
	DB, err = pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v \n", err)
		os.Exit(1)
	}

	defer DB.Close(context.Background())

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatal("Error in connecting DB")
		os.Exit(1)
	}

	log.Printf("Connected to database tasks")
}
