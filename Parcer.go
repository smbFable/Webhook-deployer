package main

import "encoding/json"

type Payload struct {
	login   string `json:"login"`
	branch  string `json:"branch"`
	message string `json:"message"`
}

func Parcer(bd []byte) (*Payload, error) {
	jsondata := &Payload{}

	err := json.Unmarshal(bd, *jsondata)
	if err != nil {
		return nil, err
	}
	return jsondata, nil
}
