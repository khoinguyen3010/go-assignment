package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Serve and listen for backend application
	app := echo.New()

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

	// Middlewares
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} | method: ${method} | endpoint: ${uri} | status: ${status} | agent: ${user_agent}\n",
	}))
	// app.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey:  []byte("wqGyEBBfPK9w3Lxw"),
	// 	TokenLookup: "query:token",
	// }))
	app.Use(middleware.Recover())

	// Routes
	app.GET("/", func(ctx echo.Context) error {
		var res struct {
			Message string `json:"message"`
		}
		res.Message = "Hello, world!"
		return ctx.JSON(http.StatusOK, res)
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
