package routers

import (
	"github.com/gin-gonic/gin"

	"goframe/app/http/controllers"
)

// RegisterAPIRoutes 注册api相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	userController := &controllers.UserController{}
	errorController := &controllers.ErrorController{}
	groupV1 := r.Group("/api/v1")

	// 简单案例
	groupV1.POST("/user_info", userController.UserInfo)

	// error test demo
	groupV1.GET("/test_error_demo1", errorController.Demo1)

}
