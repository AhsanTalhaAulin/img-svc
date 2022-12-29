package svc

import (
	"encoding/json"
	"img-svc/aws"
	"img-svc/db"
	"img-svc/domain"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo"
)

func SearchImg(c echo.Context) error {
	log.Println("Search Request Received")

	lat, err := strconv.ParseFloat(c.FormValue("lat"), 32)
	if err != nil {
		log.Println("Invalid latitude")
		return c.String(http.StatusBadRequest, "Invalid latitude")
	}

	lon, err := strconv.ParseFloat(c.FormValue("lon"), 32)
	if err != nil {
		log.Println("Invalid longitude")
		return c.String(http.StatusBadRequest, "Invalid longitude")
	}

	radius, err := strconv.ParseFloat(c.FormValue("radius"), 32)
	if err != nil {
		log.Println("Invalid radius")
		return c.String(http.StatusBadRequest, "Invalid radius")
	}

	unit := c.FormValue("radius_unit")
	// if !(unit == "km" || unit == "m") {
	// 	log.Println("Invalid radius unit")
	// 	return c.String(http.StatusBadRequest, "Invalid radius unit")
	// }

	if unit == "m" {
		radius = radius / 1000
	}

	imageList, _ := db.GetImageList(lat, lon, radius)

	presignedUrlList := getPresignedUrlList(imageList)

	result, _ := json.Marshal(presignedUrlList)

	log.Println("Search Request Served")
	return c.String(http.StatusOK, string(result))
}

func getPresignedUrlList(images []domain.Image) []string {

	var presignedUrlList []string

	for key := range images {
		presignedUrl, _ := aws.GetPresignedUrl(images[key].Uuid + filepath.Ext(images[key].Name))
		presignedUrlList = append(presignedUrlList, presignedUrl)
	}

	return presignedUrlList

}
