package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/khoinguyen3010/go-assignment/internal/services"
	postgres "github.com/khoinguyen3010/go-assignment/internal/storages/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "manabie_postgres"
	password = "thisisatest"
	dbname   = "manabie_postgres"
)

func main() {
	// Connect with Postgres database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	fmt.Print(db)
	if err != nil {
		log.Fatal("Cannot connect to Postgres database", err)
	}

	// Serve and listen for backend application
	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &postgres.Postgres{
			DB: db,
		},
	})
}
