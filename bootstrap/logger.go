package bootstrap

import (
	"goframe/infra/conf"
	"goframe/infra/logger"
	"sync"
)

var logOnce sync.Once

// SetupLogger 初始化 Logger
func SetupLogger() {
	logOnce.Do(func() {
		logger.InitLogger(
			conf.GetString("log.filename"),
			conf.GetInt("log.max_size"),
			conf.GetInt("log.max_backup"),
			conf.GetInt("log.max_age"),
			conf.GetBool("log.compress"),
			conf.GetString("log.type"),
			conf.GetString("log.level"),
			conf.GetBool("app.debug"),
		)
	})
}
