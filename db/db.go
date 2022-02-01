package db

import (
	"database/sql"
	"fmt"

	"github.com/djmin43/wedding-invitation-back-end/util"
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

func ConnectToDB() *sql.DB {
	azureHost := "wedding.postgres.database.azure.com"
	azureDatabase := "postgres"
	azureUser := "mindongjoon"
	azurePassword := "Doremi!!"
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", azureHost, azureUser, azurePassword, azureDatabase)
	db, err := sql.Open("postgres", connectionString)
	util.CheckError(err)
	err = db.Ping()
	util.CheckError(err)
	fmt.Println("connection successful nice!")
	DB = db
	return db
}


