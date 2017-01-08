package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Set up our basic static file server.
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println(http.ListenAndServe(":8080", nil))
}
