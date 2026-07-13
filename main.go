package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/PWS", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Привет, localhost!")
	})

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		return
	}
}
