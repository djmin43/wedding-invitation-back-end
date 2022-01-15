package main

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

// func getPassword() string {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	password := os.Getenv("PASSWORD")
// 	return password
// }

func connectToDB() *sql.DB {
	azureHost := "marriage-invitation.postgres.database.azure.com"
	azureDatabase := "postgres"
	azureUser := "mindongjoon"
	azurePassword := "doremi123!"
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", azureHost, azureUser, azurePassword, azureDatabase)
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("connection successful!")
	DB = db
	return db
}
