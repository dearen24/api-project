package db

import (
    "database/sql"
    "fmt"
    "log"
    "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() *sql.DB {
    // Load credentials from env vars
    cfg := mysql.NewConfig()
    cfg.User = "appuser"
    cfg.Passwd = "apppassword"
    cfg.Net = "tcp"
    cfg.Addr = "127.0.0.1:3306"
    cfg.DBName = "myapp"

    var err error
    DB, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal("Error opening DB:", err)
    }

    // Test connection
    if err := DB.Ping(); err != nil {
        log.Fatal("Error connecting to DB:", err)
    }

    fmt.Println("✅ Connected to MySQL")
	return DB
}
