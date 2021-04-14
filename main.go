package main

import (
	"net/http"

	"github.com/mrpiggy97/GoHandlers/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	handlers.Serve()
}
