package svc

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	"img-svc/aws"
)

func GetImg(c echo.Context) error {

	name := c.QueryParam("name")

	log.Println("Image name: ", name)

	urlStr, err := aws.GetPresignedUrl(name)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, urlStr)
	}

	return c.String(http.StatusOK, urlStr)
}
