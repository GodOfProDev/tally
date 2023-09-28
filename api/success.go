package main

import "encoding/json"

type APISuccessData struct {
	Status int `json:"-"`
	Data   any `json:"-"`
}

type APISuccessResponse struct {
	Status int    `json:"-"`
	Msg    string `json:"message"`
}

func (e APISuccessData) Error() string {
	jsonBytes, _ := json.Marshal(e.Data)
	return string(jsonBytes)
}

func (e APISuccessResponse) Error() string {
	return e.Msg
}
