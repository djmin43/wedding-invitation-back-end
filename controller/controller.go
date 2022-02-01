package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/djmin43/wedding-invitation-back-end/db"
	"github.com/djmin43/wedding-invitation-back-end/model"
	"github.com/djmin43/wedding-invitation-back-end/util"
)



func GetBlogs(w http.ResponseWriter, r *http.Request) {
	var id string
	var user string
	var body string
	var avatarColor string
	var createdt string
	var blogList []model.Blog

	sql_statement := `SELECT * from wedding."blogs"`
	rows, err := db.DB.Query(sql_statement)
	util.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		switch err := rows.Scan(&id, &user, &body, &createdt, &avatarColor); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			p := model.Blog{
				Id:          id,
				User:        user,
				Body:        body,
				AvatarColor: avatarColor,
				CreateDt:    createdt,
			}
			blogList = append(blogList, p)
		default:
			util.CheckError(err)
		}
	}
	jsonResp, err := json.Marshal(blogList)
	w.WriteHeader(http.StatusCreated)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
}

func AddNewPost(w http.ResponseWriter, r *http.Request) {
	var b model.Blog
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(b.AvatarColor)
	sql_statement := fmt.Sprintf(`INSERT INTO wedding.blogs (id, body, "user", createdt, avatar_color) VALUES('%s', '%s', '%s', now(), '%s');`, b.Id, b.User, b.Body, b.AvatarColor)
	defer r.Body.Close()
	rows, err := db.DB.Query(sql_statement)
	util.CheckError(err)
	defer rows.Close()
}
