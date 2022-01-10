package main

import (
	// "database/sql"
	// "fmt"

	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	// Initialize connection constants.
	HOST     = "marriage-invitation.postgres.database.azure.com"
	DATABASE = "postgres"
	USER     = "mindongjoon"
	PASSWORD = "<server_admin_password>"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getPassword() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	return port
}

func main() {
	dbPassword := getPassword()
	fmt.Println(dbPassword)
}
