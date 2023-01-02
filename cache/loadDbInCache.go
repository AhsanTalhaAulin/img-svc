package cache

import (
	"img-svc/conn"
	"img-svc/domain"
)

func LoadDbInCache() {

	var images []domain.Image

	conn.Client.Db.Find(&images)

	for index := range images {
		SaveInCache(images[index])
	}

}
