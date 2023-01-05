package db

import (
	"img-svc/conn"
	"img-svc/domain"
	"log"
	"time"
)

func GetImagesByTimestamp(timestamp string) ([]domain.Image, error) {

	imgtime, err := time.Parse(domain.TimeLayout, timestamp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("request timestamp : %v", imgtime)

	startTime := imgtime.Add(-5 * time.Minute)
	endTime := imgtime.Add(5 * time.Minute)

	var images []domain.Image

	conn.Client.Db.Where("created_at between ? and ?", startTime.Format(domain.TimeLayout), endTime.Format(domain.TimeLayout)).Find(&images)

	return images, nil
}
