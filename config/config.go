package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// 環境変数
var (
	SessionKey = makeSessionKey()
	SQLEnv     = makeSQLEnv()
	ExecuteDir = getRootPath()
)

func getRootPath() string {
	exe, _ := os.Executable()
	path := filepath.Dir(exe)
	return path
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func makeSessionKey() string {
	envLoad()
	return os.Getenv("SESSION_KEY")
}

func makeSQLEnv() string {
	envLoad()
	return os.Getenv("SQL_ENV")
}
