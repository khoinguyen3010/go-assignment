package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/khoinguyen3010/go-assignment/router"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Invoke Instance and Configure Application Middlewares
	app := router.New()

	// Connect with Postgres database
	db_port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Error(err)
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
			log.Info("Postgres DB connected...")
		}
	}

	// Run and Serve Application
	app.Logger.Fatal(app.Start(":5050"))
}
