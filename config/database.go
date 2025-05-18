package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnDB() {
	dsn := GetDotEnv("MYSQL_URL", "")
	if dsn == "" {
		log.Fatal("MYSQL_URL is not set")
	}

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to open MYSQL", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("MYSQL unreachable:", err)
	}

	log.Println("MYSQL database connected")
}
