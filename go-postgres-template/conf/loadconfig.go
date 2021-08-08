package conf

import (
	"fmt"
	"go-postgres/logrus"
	"os"

	"github.com/joho/godotenv"
)

var (
	FRONTEND_URL string
	PORT         string
	PGCS         string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		logrus.LogIt(0, "LoadConfig", "failed to load config", err)
	}

	FRONTEND_URL = os.Getenv("FRONTEND_URL")
	SQL_DB_USER := os.Getenv("SQL_DB_USER")
	SQL_DB_PASS := os.Getenv("SQL_DB_PASS")
	SQL_DB_DB := os.Getenv("SQL_DB_DB")
	SQL_DB_HOST := os.Getenv("SQL_DB_HOST")
	SQL_DB_PORT := os.Getenv("SQL_DB_PORT")
	PORT = os.Getenv("PORT")

	PGCS = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", SQL_DB_HOST, SQL_DB_USER, SQL_DB_PASS, SQL_DB_DB, SQL_DB_PORT)
}
