package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Blog struct {
	Id          string `json:"id"`
	User        string `json:"user"`
	Body        string `json:"body"`
	AvatarColor string `json:"avatarColor"`
	CreateDt    string `json:"created"`
}

func getBlogs(w http.ResponseWriter, r *http.Request) {
	var id string
	var user string
	var body string
	var avatarColor string
	var createdt string
	var blogList []Blog

	sql_statement := `SELECT * from wedding."blogs"`
	rows, err := DB.Query(sql_statement)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		switch err := rows.Scan(&id, &user, &body, &createdt, &avatarColor); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			p := Blog{
				Id:          id,
				User:        user,
				Body:        body,
				AvatarColor: avatarColor,
				CreateDt:    createdt,
			}
			blogList = append(blogList, p)
		default:
			checkError(err)
		}
	}
	jsonResp, err := json.Marshal(blogList)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
}

func addNewPost(w http.ResponseWriter, r *http.Request) {
	var b Blog
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(b.AvatarColor)
	sql_statement := fmt.Sprintf(`INSERT INTO wedding.blogs (id, body, "user", createdt, avatar_color) VALUES('%s', '%s', '%s', now(), '%s');`, b.Id, b.Body, b.User, b.AvatarColor)
	defer r.Body.Close()
	rows, err := DB.Query(sql_statement)
	checkError(err)
	defer rows.Close()
}
