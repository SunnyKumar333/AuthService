package config

import (
	env "AuthService/config/env"
	sql "database/sql"
	"fmt"
	"time"

	mysql "github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = env.GetString("User", "root")
	cfg.Passwd = env.GetString("Passwd", "root")
	cfg.Net = "tcp"
	cfg.Addr = env.GetString("Addr", "localhost:3306")
	cfg.DBName = env.GetString("DBName", "airbnb_auth_dev")

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	PingErr := db.Ping()
	if PingErr != nil {
		return nil, PingErr
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("Connected to the database Successfully!!")
	return db, nil
}
