/**
  @author: ZYL
  @date:
  @note
*/
package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.dbname"),      // use default DB
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		zap.L().Error("connect redis failed ", zap.Error(err))
	}
	return
}

// ExpireKey 让指定的key失效
func ExpireKey(key string, expiration time.Duration) error {
	if expiration <= 0 {
		expiration = 0
	}
	return rdb.Expire(key, expiration).Err()
}

func Close() {
	_ = rdb.Close()
}
