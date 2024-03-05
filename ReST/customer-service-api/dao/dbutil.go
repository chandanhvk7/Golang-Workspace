package dao

import (
	"api/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Driver   string `json:"driver"`
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func connect() *sql.DB {
	file, err := os.Open("config.json")
	utils.CheckForError(err)

	var config Config
	json.NewDecoder(file).Decode(&config)

	keys := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_PORT"}
	for _, key := range keys {
		val := os.Getenv(key)
		if val != "" {
			switch key {
			case "DB_HOST":
				config.Hostname = val
			case "DB_USER":
				config.User = val
			case "DB_PASSWORD":
				config.Password = val
			case "DB_PORT":
				config.Port = val
			}
		}
	}
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.User, config.Password, config.Hostname, config.Port, config.Database)

	db, err := sql.Open("mysql", connStr)
	utils.CheckForError(err)

	return db
}
