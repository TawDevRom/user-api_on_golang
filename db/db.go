package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB() {
	connStr := "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=userapi sslmode=disable"

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
