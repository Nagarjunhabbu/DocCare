package models

type Doctor struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Place string `json:"place"`
	Email string `json:"email"`
}
