package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
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
		fmt.Println("Неверный метод")
	}
	_, err := io.ReadAll(r.Body)
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
	fmt.Println(mac)

}
