package handlers

import (
	//"errors"
	"fmt"
	//"strings"

	"github.com/a-h/templ"
	//"github.com/emarifer/gofiber-templ-htmx/models"
	"github.com/benjamin-jack/gofiberfoxtrot/models"
	"github.com/benjamin-jack/gofiberfoxtrot/views"
	//"github.com/emarifer/gofiber-templ-htmx/views/auth_views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	//"github.com/sujit-baniya/flash"
	//"golang.org/x/crypto/bcrypt"
)

func AuthMiddleware(c *fiber.Ctx) error {
	session, _ := store.Get(c)

	fmt.Println(session.Get("AUTH_KEY"))
	fmt.Println(session.Keys())
	//checks user session and authkeys to confirm that the user is permitted to enter
	if session.Get("AUTH_KEY") == "true" {
		return c.Next() 
	} else {
		return c.Redirect("/")
	}
	//return c.Next()	
}

func HandleViewLogin(c *fiber.Ctx) error {
	session, _ := store.Get(c)
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, _ := models.GetUserByUsername(username)
	if models.CheckPassword(user, password){
		fmt.Println("PASSWORD MATCHED")
		session.Set("AUTH_KEY", "true")
		fmt.Println(session.Get("AUTH_KEY"))
		session.Save()
	} else {
		fmt.Println("PASSWORD NOT MATCHED")
	}
	// check if user information is correct,
	// if error return "incorrect login"
	// else redirect to home with logged in
	return c.Redirect("/todos")
}

func HandleViewRegister(c *fiber.Ctx) error {
	// pull register form info
	fmt.Println("REGISTER HIT")
	fmt.Println(c.Method())
	if c.Method() == "POST" {
		email := c.FormValue("email")
		password := c.FormValue("password")
		username := c.FormValue("username")
		fmt.Printf("%s %s %s", email, password, username)
		if email == "" || password == "" || username == "" {
			//msg := "Not all fields were completed, try again"
			return c.SendString("EMPTY VALS")
		}

		_, err := models.CheckEmail(email)
		fmt.Println(err)
		if err != nil {
			//msg := "Email already taken, enter a unused email"
			return c.SendString("INVALID EMAIL | ALREADY TAKEN")
		}
		
		nUser := models.User{}
		nUser.Email, nUser.Password, nUser.Username = email, password, username
		err = models.CreateUser(nUser)
		if err != nil {
			return c.SendString("SHOULD CREATE BUT DIDN'T")
		}

		return c.SendString("USER CREATED")
		//if err != nil {
			//msg := "Couldn't create user, try again"
		//}
		// if create new user with email returns not nil then return error
		// else create user and redirect to todos or home

	}
	return c.SendString("didnt post")
}

func HandleViewHome(c *fiber.Ctx) error {
	todos := views.IndexMain()
	handler := adaptor.HTTPHandler(templ.Handler(todos))
	return handler(c)	
}

