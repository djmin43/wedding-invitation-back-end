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

func getAll(w http.ResponseWriter, r *http.Request)  {
	personList := getAllApi()
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(personList)
	fmt.Fprint(w, personList)
}

func landing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func handleRequests() {
	http.HandleFunc("/all", getAll)
	http.HandleFunc("/", landing)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func main() {
	connectToDB()
	handleRequests()
}

