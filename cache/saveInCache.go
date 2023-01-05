package cache

import (
	"context"
	"img-svc/conn"
	"img-svc/domain"
	"log"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

var img domain.Image

func SaveInCache(image domain.Image) {
	img = image
	saveLocation()
	// saveTimestamp()

}

// func saveTimestamp() {
// 	var ctx = context.Background()

// 	// log.Println(img.Created_at)

// 	var timestamp redis.Z

// 	imgtime, err := time.Parse(time.RFC3339, img.Created_at)

// 	if err != nil {
// 		log.Println(err.Error())
// 	}

// 	timestamp.Score = float64(imgtime.Unix())
// 	timestamp.Member = getName()

// 	log.Println(conn.RedisClient.Rdb.ZAdd(ctx, "imageTimestamps", timestamp).String())

// }

func saveLocation() {
	var geoLocation redis.GeoLocation
	var ctx = context.Background()

	geoLocation.Name = getName()
	geoLocation.Latitude = img.Lat
	geoLocation.Longitude = img.Lon

	log.Println(conn.RedisClient.Rdb.GeoAdd(ctx, "imageLocations", &geoLocation).String())
}

func getName() string {

	imgtime, err := time.Parse(time.RFC3339, img.Created_at)
	if err != nil {
		imgtime, err = time.Parse(domain.TimeLayout, img.Created_at)

		if err != nil {

			log.Println(err.Error(), "while saving in redis")
		}
	}

	return strconv.FormatInt(imgtime.Unix(), 10) + "::" + img.Uuid + filepath.Ext(img.Name)
}
