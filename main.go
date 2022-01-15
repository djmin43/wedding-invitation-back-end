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

func getAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)
}

func handleRequests() {
	http.HandleFunc("/all", getAll)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	connectToDB()
	personList := getAllApi()
	fmt.Println(personList)
	handleRequests()
}

