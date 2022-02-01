package view

import (
	"fmt"
	"net/http"

	"github.com/djmin43/wedding-invitation-back-end/controller"
)


func Blog(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controller.GetBlogs(w, r)	
	case "POST":
		controller.AddNewPost(w, r)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}