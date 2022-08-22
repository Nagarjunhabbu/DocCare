package models

type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AuthToken string `json:"authtoken"`
}
