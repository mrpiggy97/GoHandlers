package handlers

import (
	"fmt"
	"html"
	"net/http"
)

func HomeHandler(response http.ResponseWriter, req *http.Request) {
	var message string = "this is the first version"
	fmt.Fprintf(response, "%q", html.EscapeString(message))
}

func Serve() {
	fmt.Println("listengin at port 8080")
	http.ListenAndServe("localhost:8080", nil)
}
