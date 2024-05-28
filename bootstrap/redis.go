package bootstrap

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	g "Raven-gin/global"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     g.Cof.Redis.Host + ":" + strconv.Itoa(g.Cof.Redis.Port),
		Password: g.Cof.Redis.Password,
		DB:       g.Cof.Redis.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		g.Log.Error("Redis 连接失败", zap.Any("err", err))
		return nil
	}
	return client
}
