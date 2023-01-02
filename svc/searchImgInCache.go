package svc

import (
	"context"
	"encoding/json"
	"img-svc/aws"
	"img-svc/conn"
	"img-svc/domain"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo"
)

func SearchImgInCache(c echo.Context) error {
	log.Println("Search Request Received")

	lat, err := strconv.ParseFloat(c.FormValue("lat"), 64)
	if err != nil {
		log.Println("Invalid latitude")
		return c.String(http.StatusBadRequest, "Invalid latitude")
	}

	lon, err := strconv.ParseFloat(c.FormValue("lon"), 64)
	if err != nil {
		log.Println("Invalid longitude")
		return c.String(http.StatusBadRequest, "Invalid longitude")
	}

	radius, err := strconv.ParseFloat(c.FormValue("radius"), 64)
	if err != nil {
		log.Println("Invalid radius")
		return c.String(http.StatusBadRequest, "Invalid radius")
	}

	unit := c.FormValue("radius_unit")
	if !(unit == "km" || unit == "m") {
		log.Println("Invalid radius unit")
		return c.String(http.StatusBadRequest, "Invalid radius unit")
	}

	radiusQuery := redis.GeoRadiusQuery{
		Radius:    radius,
		Unit:      unit,
		WithDist:  true,
		WithCoord: true,
	}

	var ctx = context.Background()

	images, _ := conn.RedisClient.Rdb.GeoRadius(ctx, "imageLocations", lon, lat, &radiusQuery).Result()
	var urlList []string
	for i := range images {

		log.Println("No:", i+1, images[i].Name, " Dist:", images[i].Dist, " Lat:", images[i].Latitude, " Lon:", images[i].Longitude)

		url, err := aws.GetPresignedUrl(images[i].Name)

		if err != nil {
			log.Println(err)
			url = " "
		}
		urlList = append(urlList, url)

	}
	var response domain.SearchResponse

	response.Lat = lat
	response.Lon = lon
	response.Radius = radius
	response.UrlList = urlList

	result, _ := json.Marshal(response)

	log.Println("Search Request Served From Cache")
	return c.String(http.StatusOK, string(result))

}
