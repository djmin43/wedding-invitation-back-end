package main

import (
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> f50c4ca4fcf25027bd85fb88b821367b43edbf36
	"log"
	"net/http"

	_ "github.com/lib/pq"
)


func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

<<<<<<< HEAD
func checkError(err error) {
	if err != nil {
		panic(err)
	}
=======
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}


func getAll(w http.ResponseWriter, r *http.Request)  {
	enableCors(&w)
	personList := getBlogs()
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
>>>>>>> f50c4ca4fcf25027bd85fb88b821367b43edbf36
}

func main() {
	connectToDB()
	handleRequests()
}

