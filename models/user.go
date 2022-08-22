package models

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Place string `json:"place"`
	Email string `json:"email"`
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
