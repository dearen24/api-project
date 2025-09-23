package db

import (
	"database/sql"
	"fmt"
	"go-fiber-api/config"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
