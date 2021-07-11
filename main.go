package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Invoke Instance and Configure Application Middlewares
	app := echo.New()
	rapp := app.Group("/restricted")
	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		Claims:     &jwtCustomClaims{},
	}

	// Connect with Postgres database
	db_port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println(err)
	}
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db_host, db_port, db_user, db_password, db_name)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to Postgres database", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	} else {
		if err := sqlDB.Ping(); err != nil {
			panic(err)
		} else {
			fmt.Println("Postgres DB connected...")
		}
	}

	// Normal Routes Middlewares
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} | method: ${method} | endpoint: ${uri} | status: ${status} | agent: ${user_agent}\n",
	}))
	app.Use(middleware.Recover())

	// Restricted Routes Middlewares
	rapp.Use(middleware.JWTWithConfig(jwtConfig))

	// Routes
	app.GET("/", hello)
	app.POST("/login", login)
	rapp.GET("", restricted)

	// Run and Serve Application
	app.Logger.Fatal(app.Start(":5050"))
}

func hello(ctx echo.Context) error {
	var res struct {
		Message string `json:"message"`
	}
	res.Message = "Hello, world!"
	return ctx.JSON(http.StatusOK, res)
}

func restricted(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	var res struct {
		Message string `json:"message"`
	}
	res.Message = "Restricted success! Welcome " + name
	return ctx.JSON(http.StatusOK, res)
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func login(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if username != "Khoi Nguyen" || password != "Minh@nh0806" {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		Name:  "Khoi Nguyen",
		UUID:  "9E98C454-C7AC-4330-B2EF-983765E00547",
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
