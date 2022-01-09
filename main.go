package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello")
	http.ListenAndServe(":5000", nil)
}
