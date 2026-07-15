package main

import (
	"fmt"
	"log"
	"net/http"
)

type Content struct {
	Request  string `json:"request"`
	Xhubsign string `json:"xhubsign"`
}

func main() {
	http.HandleFunc("/PWS", JSONProcessing)

	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
		return
	}
}

func JSONProcessing(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	r.Header.Add("Content-Type", "application/json")
	for secret := range r.Header {
		fmt.Fprintf(w, secret)
	}
}
