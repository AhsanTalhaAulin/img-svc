package domain

import (
	"os"
)

var DbName = "img_db"
var DbUserName = "img_user"
var DbPass = os.Args[1]

var DbPort = "3308"

var DbMaxIdleConns = 10
var DbMaxOpenConns = 10
var DbConnMaxLifetime = 1
