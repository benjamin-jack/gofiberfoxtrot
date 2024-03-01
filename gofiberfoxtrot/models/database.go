package models

import (
	"github.com/jackc/pgx/v5"
	"log"
	"context"
)

var db *pgx.Conn

func databaseConnect(){
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

func DatabaseMigrate(){
	databaseConnect()

	statement := `CREATE TABLE IF NOT EXISTS 
	users(id SERIAL PRIMARY KEY, 
	email VARCHAR(255) NOT NULL UNIQUE, 
	password VARCHAR(255) NOT NULL, 
	username VARCHAR(64) NOT NULL);`
	
	_, err := db.Exec(context.Background(),statement)
	if err != nil {
		log.Fatal(err)
	}

	// ADD A TODOS INITIALIZATION METHOD AND FIX WITH CReATEDBY VARIABLE
}

