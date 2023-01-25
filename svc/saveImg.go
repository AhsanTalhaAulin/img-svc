package svc

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"

	"img-svc/cache"
	"img-svc/db"
	"img-svc/domain"
	"img-svc/util"
)

func SaveImg(c echo.Context) error {

	start := time.Now()
	log.Printf("Post image Request Received at %v", start)
	var img domain.Image

	img.Name = c.FormValue("name")

	// imgfile, err := c.FormFile("image")
	// if err != nil {
	// 	log.Println("Error parsing the image file.")
	// 	return c.String(http.StatusBadRequest, "Error parsing the image file.")
	// }

	// imgfileOpen, _ := imgfile.Open()
	// imgfileByte, err := ioutil.ReadAll(imgfileOpen)

	// if err != nil {
	// 	log.Println(err)
	// 	return c.String(http.StatusBadRequest, "Could not process file\n")
	// }

	img.Lat, _ = strconv.ParseFloat(c.FormValue("lat"), 64)
	img.Lon, _ = strconv.ParseFloat(c.FormValue("lon"), 64)
	img.Created_at = c.FormValue("created_at")

	err := img.Validate()

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	img.Uuid, err = util.GetUniqueId(img.Created_at)
	if err != nil {
		log.Println("Error while generating unique ID :", err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	// uploadName := img.Uuid + filepath.Ext(img.Name)

	err = db.SaveInDB(img)

	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusBadRequest, "Could not save data\n")
	}

	// err = aws.UploadtoS3(uploadName, imgfileByte)

	// if err != nil {
	// 	log.Println(err.Error())
	// 	return c.String(http.StatusBadRequest, "Could not upload image\n")
	// }

	cache.SaveInCache(img, domain.TimeLayout)

	log.Printf("Post Image Request Served. Time taken: %v", time.Since(start))
	return c.String(http.StatusOK, "Request Successful")
}
