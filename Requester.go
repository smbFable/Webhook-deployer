package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

func JSONProcessing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("Неверный метод")
		return
	}

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(w, "Ошибка чтения тела запроса: ", err)
		return
	}
	defer r.Body.Close()

	sheader := r.Header.Get("X-Hub-Signature-256")
	if sheader == "" {
		fmt.Println("Заголовок с secret пуст")
		return
	}

	WHSecret := []byte("smbFableSecret1")

	mac := hmac.New(sha256.New, WHSecret)
	mac.Write(resp)
	expMAC := mac.Sum(nil)
	secret, err := hex.DecodeString(sheader[7:])

	if subtle.ConstantTimeCompare(expMAC, secret) != 1 {
		fmt.Println("Ключи не сходятся")
	}
	fmt.Println("Ключи сходятся")
}
