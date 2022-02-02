package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/djmin43/wedding-invitation-back-end/db"
	"github.com/djmin43/wedding-invitation-back-end/view"

	_ "github.com/lib/pq"
)

func handleRequests() {
	http.HandleFunc("/blog", view.Blog)
	log.Fatal(http.ListenAndServe(":80", nil))
	fmt.Println("hello world")
}

func main() {
	db.ConnectToDB()
	handleRequests()
}
