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

func SaveInCache(img domain.Image, timeLayout string) {
	var geoLocation redis.GeoLocation
	var ctx = context.Background()

	imgtime, err := time.Parse(timeLayout, img.Created_at)
	if err != nil {
		log.Println(err.Error())
	}

	geoLocation.Name = strconv.FormatInt(imgtime.Unix(), 10) + "::" + img.Uuid + filepath.Ext(img.Name)
	geoLocation.Latitude = img.Lat
	geoLocation.Longitude = img.Lon

	log.Println(conn.RedisClient.Rdb.GeoAdd(ctx, "imageLocations", &geoLocation).String())

}
