package models

import (
	"github.com/jackc/pgx/v5"
	"log"
	"context"
)

var db *pgx.Conn

func DatabaseConnect(){
	var err error
	if db != nil { return }	
	
	databaseURL := "postgres://user123:pass123@db:5432/postgres"
	db, err = pgx.Connect(context.Background(), databaseURL)

	if err !=nil {
		log.Fatalf("*|* Unable to establish onnected to database: %s", err.Error())
	}

	log.Println("*|* Connected to Database")
	return
}

func databaseMigrate(){
	DatabaseConnect()
}

