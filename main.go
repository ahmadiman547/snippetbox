package main

import (
	"log"
	"net/http"
)

// Define home handler function which write a byte slice "Hello ..." as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Starting a server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}