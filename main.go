package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	// sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect with Postgres database
	_, err := sql.Open("postgres", "manabie_postgres")
	if err != nil {
		log.Fatal("Cannot connect to Postgres database", err)
	}

	// Serve and listen for backend application
	http.ListenAndServe(":8080", nil)

	// http.ListenAndServe(":8080", &services.ToDoService{
	// 	JWTKey: "wqGyEBBfPK9w3Lxw",
	// 	Store: &sqllite.LiteDB{
	// 		DB: db,
	// 	},
	// })
}
