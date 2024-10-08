package redis

// Package redis 工具包

import (
	"context"
	"encoding/json"
	"time"

	"go.uber.org/zap"

	goRedis "github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
)

type RedisMap struct {
	mapping map[string]*RedisClient
}

const Nil = goRedis.Nil

var (
	redisMap RedisMap
)

func init() {
	redisMap = RedisMap{
		mapping: make(map[string]*RedisClient),
	}
}

// RedisClient Redis 服务
type RedisClient struct {
	Client  *goRedis.Client
	Context context.Context
	Logger  *zap.Logger
}

// ConnectRedis 连接 redis 数据库，设置全局的 Redis 对象
func ConnectRedis(clientName string, address string, username string, password string, db int, loggerInstance *zap.Logger) {
	redisClient, err := NewClient(address, username, password, db, loggerInstance)
	if err != nil {
		panic("redis 连接错误" + err.Error())
	}
	redisMap.mapping[clientName] = redisClient
}

func Client(name string) *RedisClient {
	client, ok := redisMap.mapping[name]
	if !ok {
		return nil
	}
	return client
}

// NewClient 创建一个新的 redis 连接
func NewClient(address string, username string, password string, db int, loggerInstance *zap.Logger) (*RedisClient, error) {
	// 初始化自定的 RedisClient 实例
	rds := &RedisClient{}
	rds.Logger = loggerInstance
	// 使用默认的 context
	rds.Context = context.Background()
	// 使用 redis 库里的 NewClient 初始化连接
	rds.Client = goRedis.NewClient(&goRedis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})
	// 测试一下连接
	err := rds.Ping()
	if err != nil {
		loggerInstance.Sugar().Error(err.Error())
	}
	return rds, err
}

// Ping 用以测试 redis 连接是否正常
func (rds *RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

//常用功能函数封装-start

// Set 存储 key 对应的 value，且设置 expiration 过期时间
func (rds *RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		rds.Logger.Sugar().Error("Redis", "Set", err.Error())
		return false
	}
	return true
}

// Get 获取 key 对应的 value
func (rds *RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != goRedis.Nil {
			rds.Logger.Sugar().Error("Redis", "Get", err.Error())
		}
		return ""
	}
	return result
}

// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都返回 false
func (rds *RedisClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != goRedis.Nil {
			rds.Logger.Sugar().Error("Redis", "Has", err.Error())
		}
		return false
	}
	return true
}

// Del 删除存储在 redis 里的数据，支持多个 key 传参
func (rds *RedisClient) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		rds.Logger.Sugar().Error("Redis", "Del", err.Error())
		return false
	}
	return true
}

// FlushDB 清空当前 redis db 里的所有数据
//func (rds *RedisClient) FlushDB() bool {
//	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
//		rds.Logger.Sugar().Error("Redis", "FlushDB", err.Error())
//		return false
//	}
//	return true
//}

// Increment 当参数只有 1 个时，为 key，其值增加 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要增加的值 int64 类型。 //这个函数有点问题 todo
func (rds *RedisClient) Increment(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			rds.Logger.Sugar().Error("Redis", "Increment", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
			rds.Logger.Sugar().Error("Redis", "Increment", err.Error())
			return false
		}
	default:
		rds.Logger.Sugar().Error("Redis Increment 参数过多")
		return false
	}
	return true
}

// Decrement 当参数只有 1 个时，为 key，其值减去 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要减去的值 int64 类型。
func (rds *RedisClient) Decrement(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
			rds.Logger.Sugar().Error("Redis Decrement", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
			rds.Logger.Sugar().Error("Redis Decrement", err.Error())

			return false
		}
	default:
		rds.Logger.Sugar().Error("Redis Decrement 参数过多")
		return false
	}
	return true
}

func (rds *RedisClient) GetInterface(key string) interface{} {
	stringValue := rds.Get(key)
	var wanted interface{}
	err := json.Unmarshal([]byte(stringValue), &wanted)
	if err != nil {
		rds.Logger.Sugar().Error(err.Error())
	}
	return wanted
}

// GetObject 应该传地址，用法如下:
//
//	model := user.User{}
//	cache.GetObject("key", &model)
func (rds *RedisClient) GetObject(key string, wanted interface{}) {
	val := rds.Get(key)
	if len(val) > 0 {
		err := json.Unmarshal([]byte(val), &wanted)
		if err != nil {
			rds.Logger.Sugar().Error(err.Error())
		}
	}
}

func (rds *RedisClient) GetBool(key string) bool {
	return cast.ToBool(rds.Get(key))
}

func (rds *RedisClient) GetInt64(key string) int64 {
	return cast.ToInt64(rds.Get(key))
}

func (rds *RedisClient) GetInt32(key string) int32 {
	return cast.ToInt32(rds.Get(key))
}

func (rds *RedisClient) GetFloat64(key string) float64 {
	return cast.ToFloat64(rds.Get(key))
}

func (rds *RedisClient) GetTime(key string) time.Time {
	return cast.ToTime(rds.Get(key))
}

func (rds *RedisClient) GetDuration(key string) time.Duration {
	return cast.ToDuration(rds.Get(key))
}

func (rds *RedisClient) GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(rds.GetInterface(key))
}

func (rds *RedisClient) GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(rds.GetInterface(key))
}

// 订阅处理回调函数
type SubHandle func(goRedis.Message)

func (rds *RedisClient) Sub(channel string, handler SubHandle) error {
	defer func() {
		if err := recover(); err != nil {
			rds.Logger.Sugar().Error(err.(error).Error())
		}
	}()
	pubsub := rds.Client.Subscribe(context.Background(), channel)
	_, err := pubsub.Receive(context.Background())
	if err != nil {
		rds.Logger.Sugar().Errorf("redis_sub_err:%+v", err)
		return err
	}
	ch := pubsub.Channel()
	for msg := range ch {
		handler(*msg)
	}
	return nil
}

func (rds *RedisClient) Pub(channel, msg string) bool {
	_, err := rds.Client.Publish(context.Background(), channel, msg).Result()
	if err != nil {
		rds.Logger.Sugar().Errorf("redis_pub_err:%+v", err)
		return false
	}
	return true
}
