package models

import (
	"golang.org/x/crypto/bcrypt"
	"context"
	"errors"
	//"fmt"
)

type User struct {
	ID	 uint64	`json:"id"`
	Email	 string	`json:"email"`
	Password string	`json:"password"`
	Username string	`json:"username"`
}

func CreateUser(user User) error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}	
	stmt := `INSERT INTO users(email, password, username) VALUES($1, $2, $3)`
	//Create user in pgx
	db.Exec(context.Background(), stmt, user.Email, string(hashPass), user.Username)	
	return err
}

func CheckPassword(user User, password string) bool {
	hashPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if hashPass != nil {
		return false
	} else {
		return true
	}
}

func GetUserByUsername(email string)(User, error){
	query := `SELECT * FROM users WHERE username=$1`

	tx, err := db.Begin(context.Background())
	if err != nil {
		return User{}, err
	}
	defer tx.Rollback(context.Background())

	var user User
	err = db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Username,
	)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func GetUserById(id string) (User, error) {
	query := `SELECT * FROM users WHERE id=$1`

	tx, err := db.Begin(context.Background())
	if err != nil {
		return User{}, err
	}
	defer tx.Rollback(context.Background())
	
	var user User
	err = db.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Username,
	)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func CheckEmail(email string) (User, error) {
	var err error
	//select user from email
	query := `SELECT * FROM users WHERE email=$1`

	var user User

	err = db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Username,
	)
	if err != nil {
		return User{}, nil
	}

	return user, errors.New("User already taken")

}


