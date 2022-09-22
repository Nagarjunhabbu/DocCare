package data

import (
	"DocCare/models"
	"database/sql"
	"errors"
)

type DataPatientInfo struct {
	Sql *sql.DB
}

//func to get patient list present in DB for that particular doctor
func (t DataPatientInfo) GetPatientList(docId int) (resp []models.Patient, err error) {
	query := "select id,name,place from patient where docId=?"
	rows, err := t.Sql.Query(query, docId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t models.Patient
		rows.Scan(&t.Id, &t.Name, &t.Place)
		resp = append(resp, t)
	}
	return
}

//func to get specified patient data by passing patientId
func (d DataPatientInfo) GetPatientDetails(userId int, docId int) (models.Patient, error) {
	query := "select id,name,place from patient where id=? and docId=?"
	row := d.Sql.QueryRow(query, userId, docId)
	var t models.Patient
	err := row.Scan(&t.Id, &t.Name, &t.Place)
	if err != nil {
		return models.Patient{}, err
	}
	return t, nil
}

//func to delete patient from DB
func (d DataPatientInfo) DeletePatient(userId int, docId int) (string, error) {
	_, err := d.GetPatientDetails(userId, docId)
	if err != nil {
		return "", err
	}
	query, _ := d.Sql.Prepare("delete from patient where id=? and docId=?")
	_, err = query.Exec(userId, docId)
	if err != nil {
		return "", err
	}
	return "Patient Deleted Successfully!", nil
}

//func to create new patient in DB
func (d DataPatientInfo) CreatePatient(patient models.Patient, docId int) (models.Patient, error) {

	query, _ := d.Sql.Prepare("insert into patient(name,place,docId) values (?, ?, ?)")
	resp, err := query.Exec(patient.Name, patient.Place, docId)
	if err != nil {
		return models.Patient{}, err
	}
	id, _ := resp.LastInsertId()
	return d.GetPatientDetails(int(id), docId)
}

//func to update particular information of specified user
func (d DataPatientInfo) UpdatePatient(patient models.Patient, patientId int, docId int) (models.Patient, error) {

	if patient.Place == "" {
		query, _ := d.Sql.Prepare("update patient set name=? where id=? and docId=?")
		_, err := query.Exec(patient.Name, patientId, docId)
		if err != nil {
			return models.Patient{}, err
		}
	} else if patient.Name == "" {
		query, _ := d.Sql.Prepare("update patient set place=? where id=? and docId=?")
		_, err := query.Exec(patient.Place, patientId, docId)
		if err != nil {
			return models.Patient{}, err
		}
	} else if patient.Name != "" && patient.Place != "" {
		query, _ := d.Sql.Prepare("update patient set place=?,name =? where id=? and docId=?")
		_, err := query.Exec(patient.Place, patient.Name, patientId, docId)
		if err != nil {
			return models.Patient{}, err
		}
	} else {
		return models.Patient{}, errors.New("invalid operation")
	}
	return d.GetPatientDetails(patientId, docId)
}

// func to getPatient details by passing Patient Name
func (d DataPatientInfo) GetPatientByName(patientaName string) (models.Patient, error) {
	query := "select id from patient where name=?"
	row := d.Sql.QueryRow(query, patientaName)
	var t models.Patient
	err := row.Scan(&t.Id)
	if err != nil {
		return models.Patient{}, err
	}
	return t, nil

}
