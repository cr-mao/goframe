// 项目启动初始化依赖
package bootstrap

import (
	"math/rand"
	"time"

	"goframe/app/errcode"
	"goframe/infra/conf"
)

// http服务初始化
func HttpServerBootstrap(env string) {
	// 随机数种子
	rand.New(rand.NewSource((time.Now().UnixNano())))
	// 配置初始化，依赖命令行 --env 参数
	conf.InitConfig(env)
	//全局设置时区
	var cstZone, _ = time.LoadLocation(conf.GetString("app.timezone"))
	time.Local = cstZone
	// 初始化 Logger
	SetupLogger()
	// 注册errcode
	errcode.RegisterCode()
	// 初始化数据库
	SetupDB()
	//初始化 Redis
	//SetupRedis()

}
