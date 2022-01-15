package main

import (
	"database/sql"
	"fmt"
)

func getAllApi () {
	var id string
	var name string
	var email string

	sql_statement := `SELECT * from main."user"`
	rows, err := DB.Query(sql_statement)
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