package svc

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/onokonem/sillyQueueServer/timeuuid"

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

	imgfileOpen, _ := imgfile.Open()
	imgfileByte, err := ioutil.ReadAll(imgfileOpen)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Could not process file\n")
	}

	img.Lat, _ = strconv.ParseFloat(c.FormValue("lat"), 64)
	img.Lon, _ = strconv.ParseFloat(c.FormValue("lon"), 64)
	img.Created_at = c.FormValue("created_at")

	err = img.Validate()

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	time, _ := time.Parse(domain.TimeLayout, img.Created_at)

	img.Uuid = timeuuid.UUIDFromTime(time).String()
	uploadName := img.Uuid + filepath.Ext(img.Name)

	err = db.SaveInDB(img)

	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusBadRequest, "Could not save data\n")
	}

	err = aws.UploadtoS3(uploadName, imgfileByte)

	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusBadRequest, "Could not upload image\n")
	}

	cache.SaveInCache(img, domain.TimeLayout)

	log.Println("Post Request Served")
	return c.String(http.StatusOK, "Request Successful")
}
