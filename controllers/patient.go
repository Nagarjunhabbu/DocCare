package controllers

import (
	"DocCare/models"
	"DocCare/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PatientController struct {
	Service service.ServicePatient
}

// func to getPatientList from DB
func (h PatientController) GetPatientList(c echo.Context) error {
	docId, _ := strconv.Atoi(c.Request().Header.Get("id"))

	val, err := h.Service.GetPatientList(docId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, "Patient Not Found!")
		}
		return c.JSON(http.StatusInternalServerError, "Some error Occurred! try after Sometimes")
	}
	var r models.Response
	r.Code = http.StatusOK
	r.Status = "Success"
	r.Data = val
	return c.JSON(200, r)
}

// Controller func to get particular Patientdetails by Id
func (h PatientController) GetPatientDetails(c echo.Context) error {
	docId, _ := strconv.Atoi(c.Request().Header.Get("id"))
	var patientId int
	pId := c.Param("id")
	if pId != "" {
		b, err := strconv.Atoi(pId)
		if err == nil {
			patientId = b
		}
	}
	val, err := h.Service.GetPatientDetails(patientId, docId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, "Patient Not Found!")
		}
		return c.JSON(http.StatusInternalServerError, "Some error Occurred! try after Sometimes")
	}
	var r models.Response
	r.Code = http.StatusOK
	r.Status = "Success"
	r.Data = val
	return c.JSON(200, r)
}

// Controller func to delete Patient from DB
func (h PatientController) DeletePatient(c echo.Context) error {
	docId, _ := strconv.Atoi(c.Request().Header.Get("id"))
	var patientId int
	uId := c.Param("id")
	if uId != "" {
		b, err := strconv.Atoi(uId)
		if err == nil {
			patientId = b
		}
	}
	res, err := h.Service.DeletePatient(patientId, docId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, "Access denied")
		}
		return c.JSON(http.StatusInternalServerError, "Some error Occurred! try after Sometimes")
	}
	var r models.Response
	r.Code = http.StatusOK
	r.Status = "Success"
	r.Data = res
	return c.JSON(200, r)
}

// controller func to create new Patient in DB
func (h PatientController) CreatePatient(c echo.Context) error {
	docId, _ := strconv.Atoi(c.Request().Header.Get("id"))
	var patient models.Patient
	if err := c.Bind(&patient); err != nil {
		return err
	}
	val, _ := h.Service.GetPatientByName(patient.Name)
	if val.Id <= 0 {
		res, err := h.Service.CreatePatient(patient, docId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Some error Occurred! try after Sometimes")
		}
		var r models.Response
		r.Code = http.StatusOK
		r.Status = "Success"
		r.Data = res
		return c.JSON(200, r)
	}
	return c.JSON(http.StatusInternalServerError, "Patient Already Exists!")

}

// controller func to Update particular Patient details in DB
func (h PatientController) UpdatePatient(c echo.Context) error {
	docId, _ := strconv.Atoi(c.Request().Header.Get("id"))
	var patientId int
	pId := c.Param("id")
	if pId != "" {
		b, err := strconv.Atoi(pId)
		if err == nil {
			patientId = b
		}
	}
	var user models.Patient
	if err := c.Bind(&user); err != nil {
		return err
	}
	res, err := h.Service.UpdatePatient(user, patientId, docId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, "Patient Not Found!")
		}
		return c.JSON(http.StatusInternalServerError, "Some error Occurred! try after Sometimes")
	}
	var r models.Response
	r.Code = http.StatusOK
	r.Status = "Success"
	r.Data = res
	return c.JSON(200, r)
}
