package redisBus

import (
	"os"

	"github.com/go-redis/redis"
)

var RedisAddress = os.Getenv("REDIS_ADDRESS")
var RedisPort = os.Getenv("REDIS_PORT")
var Client = &redis.Client{}
