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
	if r.Method != "POST" {
		fmt.Fprintf(w, "Ошибка метода")
		return
	}
	fmt.Println(r.Method)
	r.Header.Add("Content-Type", "application/json")
	for key, value := range r.Header {
		if key == "X-Hub-Signature-256" {
			fmt.Println(value)
		}
	}
}
