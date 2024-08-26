package database

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    dsn := "host=localhost user=postgres password=password dbname=user_info port=5432 sslmode=disable"
    var err error 
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("Failed to ping the database:", err)
    }
	log.Println("Database connection successful")
    createTable()
}

func createTable() {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        username VARCHAR(255) PRIMARY KEY,
        password VARCHAR(255) NOT NULL,
        active BOOLEAN
    );
    `
    _, err := DB.Exec(query)
    if err != nil {
        log.Fatal("Failed to create users table:", err)
    }
}