package main

import (
	"fmt"
	"net/http"
	"github.com/thiruthanikaiarasu/web/pkg/handlers"
)

const portNumber = ":8080"



// main is the main application function 
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application at %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}