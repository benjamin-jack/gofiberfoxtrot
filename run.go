package main

import (
	"context"
	"fmt"
	"os"
	"log"
	//"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
	"github.com/jackc/pgx/v5"
	//"github.com/gofiber/storage/postgres/v3"	
)
func main() {
	databaseURL := "postgres://benjamin:honeyrose@localhost:5432/fiber"
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)

	engine := handlebars.New("./views",".hbs")
	
	app := fiber.New(fiber.Config{
	    Views: engine,
	})

	app.Static("/","./")

	app.Static("/","./scripts")

//	database := postgres.New(postgres.Config{
//	Database: "fiber",
//	Username: "benjamin",
//	Password: "honeyrose",
//	})
//	database := postgres.New()
//    	var s string = "Benjamin"
//
//	sb := []byte(s)
//
//	database.Set("name", sb, 0)
//
//	var x, _ = database.Get("name")
//	var y string = string(x)

    app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello, World 👋!")
    		return c.Render("index",
			fiber.Map{"Title": "Hello, World!",})
	})
	
	app.Get("/get", func(c *fiber.Ctx) error {
  		// return c.Render("results", fiber.Map{"Results", "TEST"	})
		return c.SendString("wait")
	})

    log.Fatal(app.Listen(":3000"))
}
