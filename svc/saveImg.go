package svc

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"img-svc/aws"
	"img-svc/cache"
	"img-svc/db"
	"img-svc/domain"
)

func SaveImg(c echo.Context) error {

	// ---------------------------------------------------------------- getting data
	log.Println("Post Request Received")
	var img domain.Image

	img.Name = c.FormValue("name")

	imgfile, err := c.FormFile("image")
	if err != nil {
		log.Println("Error parsing the image file.")
		return c.String(http.StatusBadRequest, "Error parsing the image file.")
	}

	img.Lat, err = strconv.ParseFloat(c.FormValue("lat"), 64)
	if err != nil {
		log.Println("Invalid latitude")
		return c.String(http.StatusBadRequest, "Invalid latitude")
	}

	img.Lon, err = strconv.ParseFloat(c.FormValue("lon"), 64)
	if err != nil {
		log.Println("Invalid longitude")
		return c.String(http.StatusBadRequest, "Invalid longitude")
	}

	img.Created_at = c.FormValue("created_at")

	imgfileOpen, _ := imgfile.Open()
	imgfileByte, err := ioutil.ReadAll(imgfileOpen)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Could not process file\n")
	}

	img.Uuid = uuid.New().String()
	uploadName := img.Uuid + filepath.Ext(img.Name)

	err = db.SaveInDB(img)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Could not save data\n")
	}

	err = aws.UploadtoS3(uploadName, imgfileByte)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Could not upload image\n")
	}

	cache.SaveInCache(img, domain.TimeLayout)

	log.Println("Post Request Served")
	return c.String(http.StatusOK, "Request Successful")
}
