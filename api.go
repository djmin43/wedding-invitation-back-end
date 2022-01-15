package main

import (
	"database/sql"
	"fmt"
)

type person struct {
	id string
	name string
	email string
}

func getAllApi() []person {
	var id string
	var name string
	var email string

	var personList []person

	sql_statement := `SELECT * from main."user"`
	rows, err := DB.Query(sql_statement)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		switch err := rows.Scan(&id, &name, &email); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			p := person{
				id: id,
				name: name,
				email: email,
			}
			personList = append(personList, p)
			// fmt.Printf("Data row = (%s, %s, %s)\n", id, name, email)
		default:
			checkError(err)
		}
	}
	return personList
}