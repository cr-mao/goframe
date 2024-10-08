// Package config 负责配置信息
package conf

import (
	"fmt"
	"os"

	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper" // 自定义包名，避免与内置 viper 实例冲突

	"goframe/global/config"
	"goframe/infra/helpers"
)

// viper 库实例
var viper *viperlib.Viper

func init() {
	// 1. 初始化 Viper 库
	viper = viperlib.New()
	// 2. 配置类型，支持 "json", "toml", "yaml", "yml", "properties",
	//             "props", "prop", "env", "dotenv"
	viper.SetConfigType("yaml")
	// 3. 环境变量配置文件查找的路径，相对于 main.go
	viper.AddConfigPath(".")
}

// InitConfig 初始化配置信息，完成对环境变量以及 config 信息的加载
func InitConfig(env string) {
	// 1. 根据环境加载对应的配置文件
	loadEnv(env)
}

func loadEnv(envPrefix string) {
	// 默认加载 local.config.yaml 文件，如果有传参 --env=name 的话，加载 name.config.yaml 文件
	envPath := "local.config.yaml"
	if len(envPrefix) > 0 {
		filepath := envPrefix + ".config.yaml"
		if _, err := os.Stat(filepath); err == nil {
			envPath = filepath
		}
	}
	// 加载 配置文件
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 配置反序列化到全局变量中
	err := viper.Unmarshal(&config.AppConfig)
	if err != nil {
		fmt.Println(err)
	}

	// 监控 配置文件，变更时重新加载
	viper.WatchConfig()
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

// Get 获取配置项
// 第一个参数 path 允许使用点式获取，如：app.name
// 第二个参数允许传参默认值
func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	// config 或者环境变量不存在的情况
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

func GetInt32(path string, defaultValue ...interface{}) int32 {
	return cast.ToInt32(internalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}

func GetStrings(path string, defaultValue ...[]string) []string {
	value := viper.GetStringSlice(path)
	if len(value) > 0 {
		return value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return nil
}
