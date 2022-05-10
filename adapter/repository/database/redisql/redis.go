package redisql

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"time"
)

type Redis struct {
	Client *redis.Client
}

// NewRedisClient Creating a new redis client.
func (r *Redis) NewRedisClient() {
	r.Client = redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS_URL"),
		Password:    os.Getenv("REDIS_PASSWORD"),
		DialTimeout: time.Second * 20,
		DB:          0,
	})
}

// PingRedis Checking if the redis client is connected to the redis server.
func (r *Redis) PingRedis() {
	pong, err := r.Client.Ping(context.Background()).Result()
	fmt.Println(pong, err)
}

//RemoveRedisKey deletes a key from the db
func (r *Redis) RemoveRedisKey(key string) {
	err := r.Client.Del(context.Background(), key).Err()
	if err != nil {
		log.Println(err)
	}
}

//ValidateRedisKey checks if a key exist and returns its value
func (r *Redis) ValidateRedisKey(key string) (valid bool, value interface{}) {
	value, err := r.Client.Get(context.Background(), key).Result()
	if err != nil {
		log.Println(err)
		return false, nil
	}
	return true, value
}

//SetRedisKey set a redis key and value to the application redis instance
func (r *Redis) SetRedisKey(key string, value interface{}, expiration time.Duration) (valid bool, result interface{}) {
	result, err := r.Client.Set(context.Background(), key, value, expiration).Result()

	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	return true, result
}
