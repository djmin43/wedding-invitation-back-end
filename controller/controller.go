package controller

import (
	"database/sql"
	"fmt"
	"net/http"
)

type BaseHandler struct {
	db *sql.DB
}

func newBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler {
		db: db,
	}
}

func (h *BaseHandler) getAllApi (w http.ResponseWriter, r *http.Request) {
	var id string
	var name string
	var email string

	sql_statement := `SELECT * from main."user"`
	db := BaseHandler.db
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


func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
