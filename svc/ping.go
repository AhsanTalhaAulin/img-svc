package svc

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Ping(c echo.Context) error {

	log.Println("Pong ::")

	return c.String(http.StatusOK, "Pong")
}
