package cache

import (
	"context"
	"img-svc/conn"
	"img-svc/domain"
	"log"
	"path/filepath"

	"github.com/go-redis/redis/v9"
)

func SaveInCache(img domain.Image, timeLayout string) {
	var geoLocation redis.GeoLocation
	var ctx = context.Background()

	geoLocation.Name = img.Uuid + filepath.Ext(img.Name)
	geoLocation.Latitude = img.Lat
	geoLocation.Longitude = img.Lon

	log.Println(conn.RedisClient.Rdb.GeoAdd(ctx, "imageLocations", &geoLocation).String())

}
