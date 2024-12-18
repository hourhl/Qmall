package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func Init() {
	//RedisClient = redis.NewClient(&redis.Options{
	//	Addr:     conf.GetConf().Redis.Address,
	//	Username: conf.GetConf().Redis.Username,
	//	Password: conf.GetConf().Redis.Password,
	//	DB:       conf.GetConf().Redis.DB,
	//})
	// unit test
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Username: "",
		Password: "",
		DB:       0,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
