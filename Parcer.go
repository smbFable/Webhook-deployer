package main

import (
	"encoding/json"
	"fmt"
)

type Payload struct {
	Branch string `json:"ref"`
	Before string `json:"before"`
	After  string `json:"after"`
}

func Parcer(bd []byte) (*Payload, error) {
	jsondata := &Payload{}

	err := json.Unmarshal(bd, &jsondata)
	if err != nil {
		return &Payload{Branch: "Not rnown", Before: "nil", After: "nil"}, err
	}
	return jsondata, nil
}

func (pl Payload) Print() {
	fmt.Println(pl.Branch)
	fmt.Print(pl.Before, "  ->  ")
	fmt.Println(pl.After)
}
