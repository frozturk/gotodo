package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	_ "github.com/joho/godotenv/autoload"
)

var client *redis.Client

func init() {
	dsn := os.Getenv("REDISDSN")
	fmt.Println(dsn)
	client = redis.NewClient(&redis.Options{Addr: dsn})
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("REDIS CONNECTED")
}
