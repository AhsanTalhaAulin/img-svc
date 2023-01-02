package db

import (
	"img-svc/conn"
	"img-svc/domain"
	"log"
	"math"
)

const RadiusOfEarth = 6371.1

func GetImageList(lat float64, lon float64, radius float64) ([]domain.Image, error) {
	var images []domain.Image

	var result []domain.Image

	maxLat, minLat := getLatRange(lat, radius)
	maxLon, minLon := getLonRange(lon, radius)

	conn.Client.Db.Where("lat between ? and ?  AND lon between ? and ?", minLat, maxLat, minLon, maxLon).Find(&images)
	// log.Println(images)

	for key := range images {
		distance := calcDistance(lat, lon, images[key].Lat, images[key].Lon)

		if distance < radius {
			log.Println(images[key], " Distance : ", distance, " valid")
			result = append(result, images[key])
		} else {
			log.Println(images[key], " Distance : ", distance, " invalid")

		}

	}

	return result, nil
}

func getLonRange(lon float64, radius float64) (float64, float64) {
	maxLon, minLon := getRange(lon, radius)

	if maxLon > 180.0 {
		maxLon = 0.0 - (180.0 - (maxLon - 180.0))

		maxLon, minLon = minLon, maxLon
	}
	return maxLon, minLon
}
func getLatRange(lat float64, radius float64) (float64, float64) {
	maxLat, minLat := getRange(lat, radius)

	if maxLat > 90.0 {
		maxLat = 90.0 - (maxLat - 90.0)

		maxLat, minLat = minLat, maxLat
	}
	return maxLat, minLat
}

func getRange(value float64, radius float64) (float64, float64) {
	maxValue := value + (radius * (1 / ((2 * math.Pi / 360) * RadiusOfEarth)))
	minValue := value - (radius * (1 / ((2 * math.Pi / 360) * RadiusOfEarth)))
	return maxValue, minValue
}

func toRadian(value float64) float64 {
	return value / (180 / math.Pi)
}

func calcDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {

	lat1 = toRadian(lat1)
	lon1 = toRadian(lon1)
	lat2 = toRadian(lat2)
	lon2 = toRadian(lon2)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	distance := math.Pow(math.Sin(diffLat/2), 2) + math.Pow(math.Sin(diffLon/2), 2)*math.Cos(lat1)*math.Cos(lat2)

	distance = 2 * math.Asin(math.Sqrt(distance))

	distance = distance * float64(RadiusOfEarth)

	return distance
}
