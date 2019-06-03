package utils

import (
	log "github.com/Sirupsen/logrus"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var client *redis.Client

// InitRedisClient init a new redis client
func InitRedisClient() error {
	client = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.Addr"),
		Password: viper.GetString("redis.Password"), // no password set
		DB:       viper.GetInt("redis.DB"),          // use default DB
	})

	_, err := client.Ping().Result()

	return err
}

// GetRedisClient get redis client
func GetRedisClient() *redis.Client {
	return client
}

// CloseRedisClient close redis client
func CloseRedisClient() {
	err := client.Close()
	if err != nil {
		log.Error(err)
	}
	log.Info("Redis close client")
}
