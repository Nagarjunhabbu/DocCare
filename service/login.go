package service

import (
	"DocCare/data"
	"DocCare/models"

	"github.com/golang-jwt/jwt"
)

type ServiceLogin struct {
	DataLayer data.DataLoginInfo
}

// Doctor Login func
func (t ServiceLogin) Login(doctor models.Login) (models.LoginResponse, error) {
	val, _ := t.DataLayer.GetDoctorByName(doctor.Name)
	if val.Id <= 0 {
		val, _ = t.DataLayer.SignUpDoctor(doctor)
	}
	res := generateToken(val)
	var resp models.LoginResponse

	resp.AuthToken = res
	return resp, nil

}

//func to generate Auth token
func generateToken(d models.Doctor) string {
	//Todo move secret to different place
	key := "mysecretKey"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"id":         d.Id,
	})
	string, _ := token.SignedString([]byte(key))
	return string
}
