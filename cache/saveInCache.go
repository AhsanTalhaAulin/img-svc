package cache

import (
	"context"
	"img-svc/conn"
	"img-svc/domain"
	"log"
	"path/filepath"

	"github.com/go-redis/redis/v9"
)

func SaveInCache(images domain.Image) {

	var geoLocation redis.GeoLocation
	var ctx = context.Background()

	geoLocation.Name = images.Uuid + filepath.Ext(images.Name)
	geoLocation.Latitude = images.Lat
	geoLocation.Longitude = images.Lon

	log.Println(conn.RedisClient.Rdb.GeoAdd(ctx, "imageLocations", &geoLocation).String())

}
