package db

import (
	"img-svc/conn"
	"img-svc/domain"
)

func GetUUID(name string) (string, error) {
	var img domain.Image

	img.Name = name

	conn.Client.Db.Last(&img)

	return img.Uuid, nil
}
