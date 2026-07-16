package main

import (
	"fmt"
	"io"
	"net/http"
)

func AcceptRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("Неверный метод", http.StatusMethodNotAllowed)
		return
	}

	_, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("шибка чтения файла", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	content := r.Header.Get("Content-Type")
	if content != "application/json" {
		fmt.Println("Неверный тип POST запроса", http.StatusNotAcceptable)
		return
	}
	fmt.Println("Проверки выполнены успешны: ", http.StatusOK)
}
