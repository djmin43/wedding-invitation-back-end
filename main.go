package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func handleRequests() {
	http.HandleFunc("/blog", blog)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func blog(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "GET":
		fmt.Println("get request")
		blogList := getBlogs()
		fmt.Fprint(w, blogList)
	case "POST":
		addNewPost(w, r)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	connectToDB()
	handleRequests()
}
