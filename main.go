package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getPassword() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	return port
}

func connectToDB() {
	azureHost := "marriage-invitation.postgres.database.azure.com"
	azureDatabase := "postgres"
	azureUser := "mindongjoon"
	azurePassword := getPassword()
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", azureHost, azureUser, azurePassword, azureDatabase)
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("connection successful!")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	connectToDB()
}
