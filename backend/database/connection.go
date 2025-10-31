package database

import (
	"database/sql"
	"fmt"
	"log"

	"redbibliotecas/backend/config"

	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error abriendo conexion:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error conectando a base de datos:", err)
	}

	return db
}
