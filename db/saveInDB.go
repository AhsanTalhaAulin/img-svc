package db

import (
	"log"

	"img-svc/conn"
	"img-svc/domain"
)

func SaveInDB(img domain.Image) error {

	result := conn.Client.Db.Create(&img)

	log.Println("Rows affected: ", result.RowsAffected)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}
