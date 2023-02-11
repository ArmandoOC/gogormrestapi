package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// declare a db object, where we can use throughout the model package
// so in blog.go, we have access to this object
// var db *sql.DB
var DB *gorm.DB

// a struct to hold all the db connection information
type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Init() {
	err := godotenv.Load("config/.env")

	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err.Error())
		return
	}

	connInfo := connection{
		Host:     os.Getenv("POSTGRES_URL"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}
	var error error
	// try to open our postgresql connection with our connection info
	DB, error = gorm.Open(postgres.Open(connToString(connInfo)), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}

}

// Take our connection struct and convert to a string for our db connection info
func connToString(info connection) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)

}
