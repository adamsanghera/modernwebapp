package redisBus

import (
	"os"

	"github.com/go-redis/redis"
)

var redisAddress = os.Getenv("REDIS_ADDRESS")
var redisPort = os.Getenv("REDIS_PORT")
var client = &redis.Client{}
