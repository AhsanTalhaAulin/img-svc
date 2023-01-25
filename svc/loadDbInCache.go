package svc

import (
	"img-svc/cache"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func LoadDbInCache(c echo.Context) error {
	start := time.Now()
	log.Println("Load Cache Request Received")
	count := cache.LoadDbInCache()
	log.Println(count, "Rows inserted in Redis. Time taken: ", time.Since(start))
	return c.String(http.StatusOK, "Cache Loaded.")
}
