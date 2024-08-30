// Package app 应用信息
package app

import (
	"time"

	"goframe/infra/conf"
)

// 本地
func IsLocal() bool {
	return conf.Get("app.env") == "local"
}

// 是否线上环境
func IsProduction() bool {
	return conf.Get("app.env") == "production"
}

// 是否开发环境
func IsTesting() bool {
	return conf.Get("app.env") == "testing"
}

// TimenowInTimezone 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	location, _ := time.LoadLocation(conf.GetString("app.timezone"))
	return time.Now().In(location)
}
