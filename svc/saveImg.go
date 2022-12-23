package svc

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"img-svc/aws"
	"img-svc/db"
	"img-svc/domain"
)

func SaveImg(c echo.Context) error {

	// ---------------------------------------------------------------- getting data
	log.Println("Post Request Received")
	var img domain.Image

	img.Name = c.FormValue("name")
	imgfile, _ := c.FormFile("image")
	imgfileOpen, _ := imgfile.Open()
	imgfileByte, err := ioutil.ReadAll(imgfileOpen)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Could not process file\n")
	}

	img.Uuid = uuid.New().String()
	uploadName := img.Uuid + filepath.Ext(img.Name)

	err = aws.UploadtoS3(uploadName, imgfileByte)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Could not upload image\n")
	}

	err = db.SaveInDB(img)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Could not save data\n")
	}
	log.Println("Post Request Served")
	return c.String(http.StatusOK, "Request Successful")
}
