package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/WomenWhoGoTokyo/shoten10-diary-application/config"
)

func NewDB() (*sql.DB, error) {
	c := config.Config
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBPassword,
		c.DBName,
	)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
