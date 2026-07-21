package main

import (
	"io"
	"net/http"
)

func AcceptRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Неверный метод", http.StatusMethodNotAllowed)
		return
	}

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	content := r.Header.Get("Content-Type")
	if content != "application/json" {
		http.Error(w, "Неверный тип запроса", http.StatusNotAcceptable)
		return
	}

	sign := r.Header.Get("X-Hub-Signature-256")

	err = Validator(resp, sign)
	if err != nil {
		http.Error(w, "Сигнатуры не совпадают", http.StatusUnprocessableEntity)
		return
	}

	pl, err := Parcer(resp)
	if err != nil {
		http.Error(w, "Ошибка чтения JSON", http.StatusNotAcceptable)
		return
	}

	err = pl.GitValid()
	if err != nil {
		http.Error(w, "Ошибка git push", http.StatusNotAcceptable)
		return
	}

	err = Runner(pl.Repository.CloneURL, pl.Repository.FullName)
	if err != nil {
		http.Error(w, "Ошибка клонирования", http.StatusBadRequest)
	}
}
