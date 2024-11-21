package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// my ids and emails are unique so i guess composite primary key is quite fitting
type Emails struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"uniqueIndex"`
}

var DB *gorm.DB

func InitPgRepository() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)

	pgSql, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// actually I could have avoided this and just type ``postgres.Open(...)``
	// but I just wanted to try
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				Conn: pgSql,
			}),
		&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	createEmailsTable()
}

func createEmailsTable() {
	DB.AutoMigrate(&Emails{})
}
