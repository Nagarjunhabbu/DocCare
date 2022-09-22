package main

import (
	"DocCare/controllers"
	"DocCare/data"
	"DocCare/middlewares"
	"DocCare/service"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbhost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv(("DB_Port"))

	//Mysql DB connection
	sqlObj, connectionError := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+dbhost+":3306)/doccare")
	if connectionError != nil && sqlObj.Ping() != nil {
		panic(fmt.Errorf("error opening database: %v", connectionError))
	}

	//Routes
	p := controllers.PatientController{
		Service: service.ServicePatient{
			DataLayer: data.DataPatientInfo{
				Sql: sqlObj,
			},
		},
	}

	l := controllers.LoginController{
		Service: service.ServiceLogin{
			DataLayer: data.DataLoginInfo{
				Sql: sqlObj,
			},
		},
	}

	// Routes for Users
	e.GET("/api/v1/patient", p.GetPatientList, middlewares.Auth{}.AuthMiddleware())
	e.GET("/api/v1/patient/:id", p.GetPatientDetails, middlewares.Auth{}.AuthMiddleware())
	e.POST("/api/v1/patient", p.CreatePatient, middlewares.Auth{}.AuthMiddleware())
	e.DELETE("/api/v1/patient/:id", p.DeletePatient, middlewares.Auth{}.AuthMiddleware())
	e.PUT("/api/v1/patient/:id", p.UpdatePatient, middlewares.Auth{}.AuthMiddleware())

	//login endpoint
	e.POST("/login", l.Login)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, "hello")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
