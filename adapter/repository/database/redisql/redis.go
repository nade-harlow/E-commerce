package redisql

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"time"
)

// It creates a new Redis client and returns it
func connectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS_ADDR"),
		Password:    os.Getenv("REDIS_PASSWORD"),
		DialTimeout: time.Second * 20,
		DB:          0,
	})

	pong, err := client.Ping(context.Background()).Result()
	fmt.Println(pong, err)

	return client
}

//RemoveRedisKey deletes a key from the db
func RemoveRedisKey(key string) {
	redis := connectRedis()

	defer redis.Close()

	err := redis.Del(context.Background(), key).Err()

	if err != nil {
		log.Println(err)
	}
}

//ValidateRedisKey checks if a key exist and returns its value
func ValidateRedisKey(key string) (valid bool, value interface{}) {
	redis := connectRedis()

	defer redis.Close()
	value, err := redis.Get(context.Background(), key).Result()

	if err != nil {
		log.Println(err)
		return false, nil
	}
	return true, value
}

//SetRedisKey set a redis key and value to the application redis instance
func SetRedisKey(key string, value interface{}, expiration time.Duration) (valid bool, result interface{}) {
	redis := connectRedis()

	result, err := redis.Set(context.Background(), key, value, expiration).Result()

	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	return true, result
}
