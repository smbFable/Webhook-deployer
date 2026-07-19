package main

import (
	"encoding/json"
	"errors"
	"strings"
)

type Payload struct {
	Branch   string `json:"ref"`
	Before   string `json:"before"`
	After    string `json:"after"`
	ReroData struct {
		Name  string `json:"name"`
		Login string `json:"login"`
	}
}

func Parcer(bd []byte) (*Payload, error) {
	jsondata := &Payload{}

	err := json.Unmarshal(bd, &jsondata)
	if err != nil {
		return &Payload{Branch: "Not known", Before: "nil", After: "nil"}, err
	}
	return jsondata, nil
}

func (pl Payload) GitValid() error {
	substr := "/main"
	if strings.Contains(pl.Branch, substr) != true {
		return errors.New("")
	}

	if strings.EqualFold(pl.Before, pl.After) != false {
		return errors.New("")
	}
	return nil
}
