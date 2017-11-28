package redisBus

import (
	"os"

	"github.com/go-redis/redis"
)

var redisAddress = os.Getenv("REDIS_ADDRESS")
var redisPort = os.Getenv("REDIS_PORT")

//Client ...
// Interface for interacting with redis
var Client = &redis.Client{}
