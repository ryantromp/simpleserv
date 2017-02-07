package main

import (
	"fmt"
	"log"
	"net/http"
)

// Alive is the handler function which responds to the "/" endpoint
func Alive(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got a request from %s\n", r.RemoteAddr)
	fmt.Fprintln(w, "ALIVE")
}

func main() {
	http.HandleFunc("/alive", Alive)
	log.Println("Starting simpleserv on 8196")
	log.Fatal(http.ListenAndServe(":8196", nil))
}
