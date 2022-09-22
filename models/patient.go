package models

type Patient struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Place string `json:"place"`
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
