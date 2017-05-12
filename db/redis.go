package db

import (
	"github.com/go-redis/redis"
	//"github.com/Zhanat87/go/helpers"
	"os"
)

/*
@link https://github.com/go-redis/redis
@link https://redis.io/commands/get
redis-cli
 */
func NewRedis() *redis.Client {
	var dsn string
	if os.Getenv("HOME") == "/root" {
	//if helpers.IsDocker() {
		dsn = "redis:6379"
	} else {
		dsn = "localhost:6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return client
}