package main

import (
	"fmt"
	"log"
	"net/http"
)

// Alive is the handler function which responds to the "/" endpoint
func Alive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ALIVE")
}

func main() {
	http.HandleFunc("/", Alive)
	log.Fatal(http.ListenAndServe(":8196", nil))
}
