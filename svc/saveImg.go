package svc

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"img-svc/aws"
	"img-svc/db"
	"img-svc/domain"
)

func SaveImg(c echo.Context) error {

	// ---------------------------------------------------------------- getting data
	log.Println("Request received")
	var img domain.Image

	img.Name = c.FormValue("name")
	imgfile, _ := c.FormFile("image")
	imgfileOpen, _ := imgfile.Open()
	imgfileByte, err := ioutil.ReadAll(imgfileOpen)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Could not process file\n")
	}

	img.Url, err = aws.UploadtoS3(img.Name, imgfileByte)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Could not upload image\n")
	}

	err = db.SaveInDB(img)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Could not save data\n")
	}

	return c.String(http.StatusOK, "Request Successful")
}
