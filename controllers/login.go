package controllers

import (
	"DocCare/models"
	"DocCare/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginController struct {
	Service service.ServiceLogin
}

//Login func
func (h LoginController) Login(c echo.Context) error {
	var login models.Login
	if err := c.Bind(&login); err != nil {
		return err
	}
	val, err := h.Service.Login(login)
	if err != nil {
		return err
	}
	var r models.Response
	r.Code = http.StatusOK
	r.Status = "Success"
	r.Data = val
	return c.JSON(200, r)
}
