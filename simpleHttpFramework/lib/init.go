package lib

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"os"
)

func Initialize() {
	loadEnv()
	database()
	logger()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func database() {
	var err error
	Db, err = sql.Open(os.Getenv("DB_TYPE"), fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_CHARSET")))
	if err != nil {
		log.Fatal(err)
	}
	Db.SetMaxOpenConns(50)
	Db.SetMaxIdleConns(10)
	Db.Ping()
}

func logger() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer Logger.Sync()
	Sugar = Logger.Sugar()
}
