package authentication

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	DoB       time.Time `json:"date_of_birth"`
}

// func (user *User) login(ctx echo.Context) error {

// }

// func (user *User) create(ctx echo.Context) error {

// }

// func (user *User) update(ctx echo.Context) error {

// }

// func (user *User) get(ctx echo.Context) error {

// }
