package svc

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"

	"img-svc/aws"
	"img-svc/db"
)

func GetImg(c echo.Context) error {

	name := c.QueryParam("name")

	fileExt := filepath.Ext(name)
	fileName := strings.TrimSuffix(name, fileExt)
	log.Println("name: ", fileName, "ext name: ", fileExt)
	uuid, err := db.GetUUID(fileName)
	log.Println("uuid: ", uuid)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "could not get uuid")
	}

	log.Println(" uuid+ fileExt: ", uuid+fileExt)
	urlStr, err := aws.GetPresignedUrl(uuid + fileExt)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, urlStr)
	}

	return c.String(http.StatusOK, urlStr)
}