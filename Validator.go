package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
)

func Validator(bd []byte, str string) error {
	mySecret := []byte("smbFableSecret1")

	xSecret, err := hex.DecodeString(str[7:])
	if err != nil {
		return err
	}

	mac := hmac.New(sha256.New, mySecret)
	mac.Write(bd)
	aprmac := mac.Sum(nil)

	if subtle.ConstantTimeCompare(xSecret, aprmac) != 1 {
		return err
	}
	return nil
}
