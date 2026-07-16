package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"io"
	"net/http"
)

func AcceptRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Неверный метод: ", http.StatusMethodNotAllowed)
		return
	}

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения файла: ", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	content := r.Header.Get("Content-Type")
	if content != "application/json" {
		http.Error(w, "Неверный тип запроса: ", http.StatusNotAcceptable)
		return
	}

	hubsecret := []byte(r.Header.Get("X-Hub-Signature-256"))
	mysecret := []byte("smbFableSecret1")

	sha := hmac.New(sha256.New, mysecret)
	sha.Write(resp)
	aprovedsha := sha.Sum(nil)

	if subtle.ConstantTimeCompare(hubsecret, aprovedsha) != 1 {
		http.Error(w, "Ключи не сходятся: ", http.StatusBadRequest)
	}
	fmt.Println("Ключи верны")
}
