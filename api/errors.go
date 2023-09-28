package main

type APIError struct {
	Status int    `json:"-"`
	Msg    string `json:"message"`
}

func (e APIError) Error() string {
	return e.Msg
}
