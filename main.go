package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/PWS", JSONProcessing)

	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
		return
	}
}

func JSONProcessing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "Ошибка метода")
		return
	}

	r.Header.Add("Content-Type", "application/json")
	x := r.Header.Get("X-Hub-Signature-256")
	fmt.Fprintf(w, x)
}
