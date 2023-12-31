package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func CreateToken(userId int, name string) string {
	var payloadParser jwtCustomClaims

	payloadParser.ID = uint(userId)
	payloadParser.Name = name
	payloadParser.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 60))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payloadParser)
	t, _ := token.SignedString([]byte("1234"))
	return t
}

func NotFoundHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		if err != nil {
			if he, ok := err.(*echo.HTTPError); ok {
				if he.Code == http.StatusNotFound {
					errorMessage := "Invalid Endpoint"
					return c.JSON(http.StatusNotFound, map[string]interface{}{
						"message": errorMessage,
					})
				}
			}

			fmt.Println("Terjadi kesalahan:", err)
		}

		return err
	}
}

func Logger(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
}
