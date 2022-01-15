package main

import (
	"database/sql"
	"fmt"
)

type Person struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func getAllApi() []Person {
	var id string
	var name string
	var email string

	var personList []Person

	sql_statement := `SELECT * from main."user"`
	rows, err := DB.Query(sql_statement)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		switch err := rows.Scan(&id, &name, &email); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			p := Person{
				Id: id,
				Name: name,
				Email: email,
			}
			personList = append(personList, p)
			// fmt.Printf("Data row = (%s, %s, %s)\n", id, name, email)
		default:
			checkError(err)
		}
	}
	return personList
}