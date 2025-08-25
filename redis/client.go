package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

//go get package github.com/go-redis/redis/v8


var Ctx = context.Background()
var Client *redis.Client
var Nil = redis.Nil

// ConnectRedis initializes the Redis client
func ConnectRedis() error {
	var err error

	// Prefer REDIS_URL (used by Render Key Value)
	if redisURL := os.Getenv("REDIS_URL"); redisURL != "" {
		var opt *redis.Options
		opt, err = redis.ParseURL(redisURL)
		if err != nil {
			return fmt.Errorf("invalid REDIS_URL: %v", err)
		}
		Client = redis.NewClient(opt)
		fmt.Println("Connecting to Redis using REDIS_URL")
	} else {
		// Fallback: loca with REDIS_ADDR + REDIS_PASSWORD
		addr := os.Getenv("REDIS_ADDR")
		if addr == "" {
			addr = "localhost:6379" // default local Redis
		}
		Client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: os.Getenv("REDIS_PASSWORD"), // "" if no password
			DB:       0,
		})
		fmt.Println("Connecting to Redis using REDIS_ADDR")
	}

	// Test the connection
	_, err = Client.Ping(Ctx).Result()
	if err != nil {
		return fmt.Errorf("could not connect to Redis: %v", err)
	}

	fmt.Println("✅ Connected to Redis")
	return nil
}