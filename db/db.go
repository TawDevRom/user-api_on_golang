package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB() {
	// connStr := "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=userapi sslmode=disable"
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal("БД не отвечает:", err)
	}

	log.Println("Подключение к PostgreSQL успешно!")

	_, err = Db.Exec(`CREATE TABLE IF NOT EXISTS users 
	(
		id SERIAL PRIMARY KEY,
		name TEXT,
		age INTEGER,
		login TEXT UNIQUE,
		password TEXT
	)`)

	if err != nil {
		log.Fatal("Ошибка создания таблиц:", err)
	}
}
