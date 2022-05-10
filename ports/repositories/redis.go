package repositories

import "time"

type Redis interface {
	PingRedis()
	RemoveRedisKey(key string)
	ValidateRedisKey(key string) (valid bool, value interface{})
	SetRedisKey(key string, value interface{}, expiration time.Duration) (valid bool, result interface{})
}
