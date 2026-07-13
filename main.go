package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/PWS", JSONProcessing)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
		return
	}
}

func JSONProcessing(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "application/json")
	fmt.Fprintln(w, "Hello, World!")
}
