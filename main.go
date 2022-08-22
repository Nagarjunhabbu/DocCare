package main

import (
	"chkdin/controllers"
	"chkdin/data"
	"chkdin/middlewares"
	"chkdin/service"
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

	//Mysql DB connection
	sqlObj, connectionError := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+dbhost+":3306)/chkdin")
	if connectionError != nil && sqlObj.Ping() != nil {
		panic(fmt.Errorf("error opening database: %v", connectionError))
	}

	//Routes
	u := controllers.UsersController{
		Service: service.ServiceUsers{
			DataLayer: data.DataUserInfo{
				Sql: sqlObj,
			},
		},
	}

	l := controllers.LoginController{
		Service: service.ServiceLogin{
			DataLayer: data.DataUserInfo{
				Sql: sqlObj,
			},
		},
	}

	// Routes for Users
	e.GET("/api/v1/user", u.GetUserList, middlewares.Auth{}.AuthMiddleware())
	e.GET("/api/v1/user/:id", u.GetUserDetails, middlewares.Auth{}.AuthMiddleware())
	e.POST("/api/v1/user", u.CreateUser, middlewares.Auth{}.AuthMiddleware())
	e.DELETE("/api/v1/user/:id", u.DeleteUser, middlewares.Auth{}.AuthMiddleware())
	e.PUT("/api/v1/user/:id", u.UpdateUser, middlewares.Auth{}.AuthMiddleware())

	//login endpoint
	e.POST("/login", l.Login)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
