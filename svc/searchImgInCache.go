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
	"strings"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo"
)

func SearchImgInCache(c echo.Context) error {
	start := time.Now()
	log.Printf("Search Request Received at %v", start)

	var searchRequest domain.SearchRequest

	if err := c.Bind(&searchRequest); err != nil {
		return c.String(http.StatusBadRequest, "could not bind request")
	}

	err := searchRequest.Validate()

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	log.Println(searchRequest)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	imagesByLocation, err := getImagesByLocation(searchRequest)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	// imagesByTimestamp, err := db.GetImagesByTimestamp(searchRequest.Timestamp)
	// log.Println(imagesByTimestamp)

	// if err != nil {
	// 	log.Println(err)
	// 	return c.String(http.StatusBadRequest, err.Error())
	// }

	urlList := getUrlListByTimestamp(imagesByLocation, searchRequest.Timestamp, 300)

	var response domain.SearchResponse

	response.Lat = searchRequest.Lat
	response.Lon = searchRequest.Lon
	response.Radius = searchRequest.Radius
	response.Unit = searchRequest.Unit
	response.Timestamp = searchRequest.Timestamp
	response.UrlList = urlList

	result, _ := json.Marshal(response)

	log.Printf("Search Request Served From Cache. Time taken: %v", time.Since(start))
	// log.Println("Search Request Served From Cache")
	return c.String(http.StatusOK, string(result))

}

func getUrlListByTimestamp(images []redis.GeoLocation, timestamp string, rangeInSeconds int64) []string {
	var urlList []string

	time, _ := time.Parse(domain.TimeLayout, timestamp)

	startTime := time.Unix() - rangeInSeconds
	endTime := time.Unix() + rangeInSeconds

	for i := range images {

		// info format is 1702417932::6f04d29e-d1e7-4648-883c-5537281c9634.png

		info := strings.Split(images[i].Name, "::")

		created_at, _ := strconv.ParseInt(info[0], 10, 64)
		name := info[1]

		if created_at >= startTime && created_at <= endTime {

			url, err := aws.GetPresignedUrl(name)

			if err != nil {
				log.Println(err)
				url = "Could not get presigned url"
			}
			urlList = append(urlList, url)

			log.Printf("%v : %v Diff: %v ---> valid\n", name, created_at, time.Unix()-created_at)

		} else {

			log.Printf("%v : %v Diff: %v ---> invalid\n", name, created_at, time.Unix()-created_at)
		}

	}

	return urlList
}

// func GetImagesByTimestamp(searchRequest domain.SearchRequest) ([]string, error) {
// 	timestamp, err := time.Parse(time.RFC3339, searchRequest.Timestamp)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	startTime := timestamp.Add(5 * time.Minute)
// 	endTime := timestamp.Add(-5 * time.Minute)

// 	var ctx = context.Background()

// 	redisTimestampQuery := conn.RedisClient.Rdb.ZRangeByScore(ctx, "imageTimestamps", startTime.Unix(), endTime.Unix())

// 	log.Println(redisTimestampQuery.String())

// 	imagesByTimestamp, err := redisTimestampQuery.Result()
// 	if err != nil {
// 		log.Println(err.Error())
// 		return nil, err
// 	}
// 	return imagesByTimestamp, nil

// }

func getImagesByLocation(searchRequest domain.SearchRequest) ([]redis.GeoLocation, error) {
	radiusQuery := redis.GeoRadiusQuery{
		Radius:    searchRequest.Radius,
		Unit:      searchRequest.Unit,
		WithDist:  true,
		WithCoord: true,
	}
	var ctx = context.Background()

	redisLocationQuery := conn.RedisClient.Rdb.GeoRadius(ctx, "imageLocations", searchRequest.Lon, searchRequest.Lat, &radiusQuery)
	// log.Println(redisLocationQuery.String())
	imagesByLocation, err := redisLocationQuery.Result()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return imagesByLocation, nil

}
