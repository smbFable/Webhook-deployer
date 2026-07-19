package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/PWS", AcceptRequest)

	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
		return
	}
	fmt.Println("Контейнер завершен")
}
