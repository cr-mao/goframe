package routers

import (
	"github.com/gin-gonic/gin"

	"goframe/app/errcode"
	"goframe/app/http/middlewares"
	"goframe/app/http/response"
	"goframe/infra/app"
	"goframe/infra/conf"
)

// 404处理
func setup404Handler(r *gin.Engine) {
	// 添加 Get 请求路路由
	r.NoRoute(func(c *gin.Context) {
		response.Write(c, 404, response.Response{
			ErrorCode: errcode.ErrRouteNotFound,
			Msg:       "找不到路由",
			Data:      nil,
		})
	})
}

// 全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),   // 请求错误 log记录
		middlewares.Recovery(), //panic   错误 拦截处理
	)
}

func NewRouter() *gin.Engine {
	if app.IsLocal() && conf.GetBool("app.debug", false) {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	// 全局中间件
	registerGlobalMiddleWare(router)
	// 404处理
	setup404Handler(router)
	//外部api
	RegisterAPIRoutes(router)
	return router
}
