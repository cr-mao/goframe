package bootstrap

import (
	"fmt"
	"strconv"
	"sync"

	"goframe/global"
	"goframe/infra/conf"
	"goframe/infra/console"
	"goframe/infra/logger"
	"goframe/infra/redis"
)

var redisOnce sync.Once

// SetupRedis 初始化 Redis
func SetupRedis() {
	// 建立 Redis 连接
	redisOnce.Do(func() {
		//专门做cache用
		cacheAddress := fmt.Sprintf("%v:%v", conf.GetString("redis.default.host"), conf.GetString("redis.default.port"))
		redis.ConnectRedis(
			global.REDIS_DEFAULT,
			cacheAddress,
			conf.GetString("redis.default.username"),
			conf.GetString("redis.default.password"),
			conf.GetInt("redis.default.database"),
			logger.Logger,
		)
		console.Success("redis cache connect success:" + cacheAddress + " db:" + strconv.Itoa(conf.GetInt("redis.cache.database")))
	})
}
