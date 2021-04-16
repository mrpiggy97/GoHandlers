package handlers

import (
	"fmt"
	"html"
	"net/http"
)

func HomeHandler(response http.ResponseWriter, req *http.Request) {
	var message string = "this is supposed to be the mayor version 1.13.9 version"
	fmt.Fprintf(response, "%q", html.EscapeString(message))
}

func Serve() {
	fmt.Println("listening at port 8080")
	http.ListenAndServe("localhost:8080", nil)
}
