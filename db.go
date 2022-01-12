package main

import (
	"database/sql"
	"fmt"
)

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

	var id string
	var name string
	var email string

	sql_statement := `SELECT * from main."user"`
	rows, err := db.Query(sql_statement)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		switch err := rows.Scan(&id, &name, &email); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%s, %s, %s)\n", id, name, email)
		default:
			checkError(err)
		}
	}

}
