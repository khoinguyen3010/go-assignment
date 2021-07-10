package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func main() {
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
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Cannot connect to Postgres database", err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("Postgres DB connected...")
	}

	// Serve and listen for backend application
	app := echo.New()

	// Middlewares
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Routes
	app.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, world!")
	})

	// Run application
	app.Logger.Fatal(app.Start(":5050"))

	// http.ListenAndServe(":5050", &services.ToDoService{
	// 	JWTKey: "wqGyEBBfPK9w3Lxw",
	// 	Store: &postgres.Postgres{
	// 		DB: db,
	// 	},
	// })
}
