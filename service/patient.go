package service

import (
	"DocCare/data"
	"DocCare/models"
)

type ServicePatient struct {
	DataLayer data.DataPatientInfo
}

func (t ServicePatient) GetPatientList(docId int) ([]models.Patient, error) {
	val, err := t.DataLayer.GetPatientList(docId)
	if err != nil {
		return nil, err
	}
	return val, nil

}

func (p ServicePatient) GetPatientDetails(userId int, docId int) (models.Patient, error) {
	val, err := p.DataLayer.GetPatientDetails(userId, docId)
	if err != nil {
		return models.Patient{}, err
	}
	return val, nil
}

func (p ServicePatient) DeletePatient(userId int, docId int) (string, error) {
	res, err := p.DataLayer.DeletePatient(userId, docId)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (p ServicePatient) CreatePatient(patient models.Patient, docId int) (models.Patient, error) {
	res, err := p.DataLayer.CreatePatient(patient, docId)
	if err != nil {
		return models.Patient{}, err
	}
	return res, nil
}

func (p ServicePatient) UpdatePatient(user models.Patient, userId int, docId int) (models.Patient, error) {
	res, err := p.DataLayer.UpdatePatient(user, userId, docId)
	if err != nil {
		return models.Patient{}, err
	}
	return res, nil
}

func (p ServicePatient) GetPatientByName(patientName string) (models.Patient, error) {
	res, err := p.DataLayer.GetPatientByName(patientName)
	if err != nil {
		return models.Patient{}, err
	}
	return res, nil
}
