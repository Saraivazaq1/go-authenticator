package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	DBSSLMode  string
	JWTSecret  string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: .env não carregado (variáveis precisam estar definidas no sistema)")
	}

	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBName = os.Getenv("DB_NAME")
	DBSSLMode = os.Getenv("DB_SSLMODE")
}

func GetDSN() string {
	return "user=" + DBUser +
		" password=" + DBPassword +
		" host=" + DBHost +
		" port=" + DBPort +
		" dbname=" + DBName +
		" sslmode=" + DBSSLMode
}