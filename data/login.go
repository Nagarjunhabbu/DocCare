package data

import (
	"DocCare/models"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type DataLoginInfo struct {
	Sql *sql.DB
}

// func to getDoctor  details by passing doctor Name
func (d DataLoginInfo) GetDoctorByName(docName string) (models.Doctor, error) {
	query := "select id from doctors where name=?"
	row := d.Sql.QueryRow(query, docName)
	var t models.Doctor
	err := row.Scan(&t.Id)
	if err != nil {
		return models.Doctor{}, err
	}
	return t, nil

}

//func to register new doctor
func (d DataLoginInfo) SignUpDoctor(doctor models.Login) (models.Doctor, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(doctor.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	password := string(hash)

	query, _ := d.Sql.Prepare("insert into doctor(name,password) values (?, ?)")
	resp, err := query.Exec(doctor.Name, password)
	if err != nil {
		return models.Doctor{}, err
	}
	id, _ := resp.LastInsertId()
	return d.GetDoctorDetails(int(id))
}

//func to get specified doctor data by passing docId
func (d DataLoginInfo) GetDoctorDetails(docId int) (models.Doctor, error) {
	query := "select id,name,place,email from doctor where id=?"
	row := d.Sql.QueryRow(query, docId)
	var t models.Doctor
	err := row.Scan(&t.Id, &t.Name, &t.Place, &t.Email)
	if err != nil {
		return models.Doctor{}, err
	}
	return t, nil
}
