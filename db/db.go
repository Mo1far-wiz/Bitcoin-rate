package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPgRepository() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	createTables()
}

func createTables() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS emails (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE
	);`

	_, err := DB.Exec(createUserTable)

	if err != nil {
		log.Panic("Could not create table.")
	}
}
