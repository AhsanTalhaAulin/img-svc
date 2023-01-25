package domain

import "os"

var RedisHost = os.Getenv("RedisHost")
var RedisPort = os.Getenv("RedisPort")
