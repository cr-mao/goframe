/**
* @Author: cr-mao
* @Desc:
**/
package config

type App struct {
	Name             string `json:"name" mapstructure:"name"`   // 应用名称
	Env              string `json:"env" mapstructure:"env"`     // 环境
	Debug            bool   `json:"debug" mapstructure:"debug"` //是否开启调试模式
	DatabaseTimeZone string `json:"database_timezone" mapstructure:"database_timezone"`
	TimeZone         string `json:"timezone"  mapstructure:"timezone"`
	HttpHost         string `json:"http_host" mapstructure:"http_host"`
	HttpPort         int    `json:"http_port" mapstructure:"http_port"`
}

type Config struct {
	App *App `json:"app" mapstructure:"app"`
}

// 应用所有配置
var AppConfig Config
