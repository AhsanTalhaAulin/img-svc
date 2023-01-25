package domain

import "os"

// var DbName = "img_db"
// var DbUserName = "img_user"

// var DbPass = "12345678"

// var DbHost = "db"
var DbHost = os.Getenv("DbHost")

var DbName = os.Getenv("DbName")

var DbUserName = os.Getenv("DbUserName")
var DbPass = os.Getenv("DbPass")

var DbPort = "3306"

var DbMaxIdleConns = 10
var DbMaxOpenConns = 10
var DbConnMaxLifetime = 1
