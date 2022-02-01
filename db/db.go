package db

import (
	"database/sql"
	"fmt"

	"github.com/djmin43/wedding-invitation-back-end/util"
)

var DB *sql.DB

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
	fmt.Println("connection successful")
	DB = db
	return db
}


