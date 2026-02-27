package inits

import (
	"ant-grid/internal/common/global"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func RedisInit() {
	var data = global.AppConf.Redis
	addr := fmt.Sprintf("%s:%d", data.Host, data.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: data.Password, // no password set
		DB:       data.DB,       // use default DB
	})

	err := rdb.Set(context.Background(), "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("redis init success")
}
