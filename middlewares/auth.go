package middlewares

import (
	"DocCare/models"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Auth struct {
}

//func to Authenticate User
func (a Auth) AuthMiddleware() echo.MiddlewareFunc {
	key := "mysecretKey"
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var r models.Response
			r.Code = 401
			r.Status = "un authorozed"
			//do the things
			auth := c.Request().Header.Get("auth")
			token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
				//Make sure that the token method conform to "SigningMethodHMAC"
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(key), nil
			})
			if err != nil {
				return c.JSON(401, r)
			}
			claims, ok := token.Claims.(jwt.MapClaims)

			if token.Valid && ok {
				id, ok := claims["id"].(float64)
				if !ok {
					return c.JSON(401, r)
				}
				c.Request().Header.Set("id", fmt.Sprint(id))
			}

			return next(c)
		}

	}
}
