package main

import (
	"database/sql"
	"fmt"
)

type Blog struct {
	Id string `json:"id"`
	User string `json:"user"`
	Body string `json:"body"`
	AvatarColor string `json:"avatar_color"`
}

func getBlogs() []Blog {
	var id string
	var user string
	var body string
	var avatar_color string
	var blogList []Blog

	sql_statement := `SELECT * from wedding."blogs"`
	rows, err := DB.Query(sql_statement)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		switch err := rows.Scan(&id, &user, &body); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			p := Blog{
				Id: id,
				User: user,
				Body: body,
				AvatarColor: avatar_color,
			}
			blogList = append(blogList, p)
		default:
			checkError(err)
		}
	}
	return blogList
}

