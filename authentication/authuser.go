package authentication

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	DoB       time.Time `json:"date_of_birth"`
}

type UserLoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	secret := os.Getenv("JWT_SECRET")

	// TODO: query from database for user instead of hard code
	if username != "Khoi Nguyen" || password != "Minh@nh0806" {
		return echo.ErrUnauthorized
	}

	claims := &JwtCustomClaims{
		Name:  "Khoi Nguyen",
		UUID:  "9E98C454-C7AC-4330-B2EF-983765E00547",
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Error(err)
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": t})
}

func Restricted(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	name := claims.Name
	var res struct {
		Message string `json:"message"`
	}
	res.Message = "Restricted success! Welcome " + name
	return ctx.JSON(http.StatusOK, res)
}

func Index(ctx echo.Context) error {
	var res struct {
		Message string `json:"message"`
	}
	res.Message = "Hello, world!"
	return ctx.JSON(http.StatusOK, res)
}

// func (user *User) login(ctx echo.Context) error {

// }

// func (user *User) create(ctx echo.Context) error {

// }

// func (user *User) update(ctx echo.Context) error {

// }

// func (user *User) get(ctx echo.Context) error {

// }
