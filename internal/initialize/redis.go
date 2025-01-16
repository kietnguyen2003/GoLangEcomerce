package initialize

import (
	"context"
	"fmt"
	"src/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Db,       // use default DB
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("redis connect failed, err:", zap.Error(err))
		panic(err)
	}

	global.Redis = rdb
	global.Logger.Info("redis connect success")

}
