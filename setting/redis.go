package setting

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func RedisClient() {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码就填空字符长即可
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		zap.L().Error("redis数据库连接错误:", zap.Error(err))
	} else {
		fmt.Println("redis连接成功")
	}

}
