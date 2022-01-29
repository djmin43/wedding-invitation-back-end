package main

import (
	"fmt"
	"io/ioutil"
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
		blogList := getBlogs()
		fmt.Fprint(w, blogList)
	case "POST":
    fmt.Println("Method : ", r.Method)
    fmt.Println("URL : ", r.URL)
    fmt.Println("Header : ", r.Header)
    b, _ := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    fmt.Println("Body : ", string(b))		

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	connectToDB()
	handleRequests()
}

