package service

import (
	"chkdin/data"
	"chkdin/models"

	"github.com/golang-jwt/jwt"
)

type ServiceLogin struct {
	DataLayer data.DataUserInfo
}

//Login func
func (t ServiceLogin) Login(user models.Login) (models.LoginResponse, error) {
	val, _ := t.DataLayer.GetUserByName(user.Name)
	if val.Id <= 0 {
		val, _ = t.DataLayer.SignUpUser(user)
	}
	res := generateToken(val)
	var resp models.LoginResponse

	resp.AuthToken = res
	return resp, nil

}

//func to generate Auth token
func generateToken(u models.User) string {
	//Todo move secret to different place
	key := "mysecretKey"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"user":       u,
	})
	string, _ := token.SignedString([]byte(key))
	return string
}
