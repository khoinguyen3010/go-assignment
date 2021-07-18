package main

import (
	"github.com/khoinguyen3010/go-assignment/models"
	"github.com/khoinguyen3010/go-assignment/router"
	"github.com/khoinguyen3010/go-assignment/storage"
	// _ "github.com/lib/pq"
)

func main() {
	// Invoke Instance and Configure Application Middlewares
	app := router.New()

	// Initializing and connecting to Database
	db, err := storage.InitDB()
	if err != nil {
		panic(err)
	}
	// Migration
	db.AutoMigrate(
		&models.User{},
		&models.Task{},
	)

	// Run and Serve Application
	app.Logger.Fatal(app.Start(":5050"))
}

// Routes call handler -> handler call struct method -> call orm to handle
