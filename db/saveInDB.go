package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"img-svc/domain"
)

func SaveInDB(img domain.Image) error {
	dsn := "root:12345678@tcp(127.0.0.1:3308)/img_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	log.Println("DB connected")

	result := db.Create(&img)

	log.Println("Rows affected: ", result.RowsAffected)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return nil
}
