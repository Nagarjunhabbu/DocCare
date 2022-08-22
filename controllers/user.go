package controllers

import (
	"chkdin/models"
	"chkdin/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UsersController struct {
	Service service.ServiceUsers
}

// func to getUserList from DB
func (h UsersController) GetUserList(c echo.Context) error {
	val, err := h.Service.GetUsersList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var r models.Response
	r.Code = http.StatusOK
	r.Status = "Success"
	r.Data = val
	return c.JSON(200, r)
}

// Controller func to get particular userdetails by Id
func (h UsersController) GetUserDetails(c echo.Context) error {
	var userId int
	uId := c.Param("id")
	if uId != "" {
		b, err := strconv.Atoi(uId)
		if err == nil {
			userId = b
		}
	}
	val, err := h.Service.GetUserDetails(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "User not found!")
	}
	var r models.Response
	r.Code = http.StatusOK
	r.Status = "Success"
	r.Data = val
	return c.JSON(200, r)
}

// Controller func to delete User from DB
func (h UsersController) DeleteUser(c echo.Context) error {
	var userId int
	uId := c.Param("id")
	if uId != "" {
		b, err := strconv.Atoi(uId)
		if err == nil {
			userId = b
		}
	}
	res, err := h.Service.DeleteUser(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "User not found!")
	}
	var r models.Response
	r.Code = http.StatusOK
	r.Status = "Success"
	r.Data = res
	return c.JSON(200, r)
}

// controller func to create new user in DB
func (h UsersController) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	res, err := h.Service.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Some error Occurred! try after Sometimes")
	}
	var r models.Response
	r.Code = http.StatusOK
	r.Status = "Success"
	r.Data = res
	return c.JSON(200, r)
}

// controller func to Update particular user details in DB
func (h UsersController) UpdateUser(c echo.Context) error {
	var userId int
	uId := c.Param("id")
	if uId != "" {
		b, err := strconv.Atoi(uId)
		if err == nil {
			userId = b
		}
	}
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	res, err := h.Service.UpdateUser(user, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Some error Occurred! try after Sometimes")
	}
	var r models.Response
	r.Code = http.StatusOK
	r.Status = "Success"
	r.Data = res
	return c.JSON(200, r)
}
