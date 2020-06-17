package database

import (
	"fmt"
	"match-schedule/pkg/configs"

	"github.com/go-redis/redis"
)

// RedisClient redis client
var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configs.EnvConfig.Redis.Host, configs.EnvConfig.Redis.Port),
		Password: configs.EnvConfig.Redis.Pwd, // no password set
		DB:       0,                           // use default DB
	})
	pong, err := RedisClient.Ping().Result()
	fmt.Println(pong, err)
}
