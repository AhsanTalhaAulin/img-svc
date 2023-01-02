package svc

import (
	"img-svc/cache"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func LoadDbInCache(c echo.Context) error {
	log.Println("Load Cache Request Received")
	cache.LoadDbInCache()
	log.Println("Load Cache Request Served")
	return c.String(http.StatusOK, "Cache Loaded")
}
