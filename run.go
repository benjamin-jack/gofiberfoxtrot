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
	fmt.Println("TESTING DATABASE CONNECTION")
	databaseURL := "postgres://user123:pass123@db:5432/postgres"
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
//	var weight int64
	fmt.Println("Name query sent")
	err = conn.QueryRow(context.Background(), "SELECT name from todo where isdone=true;").Scan(&name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("complete")
	//fmt.Println(name, weight)

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
		return c.SendString(name)
	})

    log.Fatal(app.Listen(":3000"))
}
