package svc

import (
	"context"
	"encoding/json"
	"img-svc/aws"
	"img-svc/conn"
	"img-svc/domain"
	"img-svc/util"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo"
)

func SearchImgInCache(c echo.Context) error {
	start := time.Now()
	// log.Printf("Search Request Received at %v", start)

	var searchRequest domain.SearchRequest

	if err := c.Bind(&searchRequest); err != nil {
		log.Println(err.Error())
		return c.String(http.StatusBadRequest, "could not bind request")
	}

	err := searchRequest.Validate()

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	// log.Println(searchRequest)

	images, err := getImagesByLocation(searchRequest)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	urlList, err := getUrlListByTimeStamp(images, searchRequest)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	var response domain.SearchResponse

	response.Lat = searchRequest.Lat
	response.Lon = searchRequest.Lon
	response.Radius = searchRequest.Radius
	response.Unit = searchRequest.Unit
	response.Timestamp = searchRequest.Timestamp
	response.UrlList = urlList

	result, _ := json.Marshal(response)

	log.Println(searchRequest, " Request Served. Total :", len(urlList), " Time taken: ", time.Since(start))
	return c.String(http.StatusOK, string(result))

}

func getUrlListByTimeStamp(images []string, searchRequest domain.SearchRequest) ([]string, error) {

	time, _ := time.Parse(domain.TimeLayout, searchRequest.Timestamp)
	startTime := time.Unix() - 300
	endTime := time.Unix() + 300

	var urlList []string
	for _, name := range images {

		uid := strings.TrimSuffix(name, filepath.Ext(name))

		timeStamp, err := util.GetUnixTime(uid)
		if err != nil {
			log.Println(err)
		}

		// log.Println(timeStamp)

		if timeStamp >= startTime && timeStamp <= endTime {
			url, err := aws.GetPresignedUrl(name)
			if err != nil {
				log.Println(err)
				url = err.Error()
			}
			urlList = append(urlList, url)
		}

	}

	return urlList, nil

}

func getImagesByLocation(searchRequest domain.SearchRequest) ([]string, error) {
	var ctx = context.Background()
	searchQuery := redis.GeoSearchQuery{
		Longitude:  searchRequest.Lon,
		Latitude:   searchRequest.Lat,
		Radius:     searchRequest.Radius,
		RadiusUnit: searchRequest.Unit,
	}
	images, err := conn.RedisClient.Rdb.GeoSearch(ctx, "imageLocations", &searchQuery).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// log.Println(images)

	return images, nil
}
