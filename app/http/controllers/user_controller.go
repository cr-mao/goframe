/**
* @Author: cr-mao
* @Desc: 用户控制器
**/
package controllers

import (
	"github.com/gin-gonic/gin"

	"goframe/app"
	"goframe/app/http/requests"
	"goframe/app/http/response"
)

// 用户信息返回返回,vo 对应 直接在controller 写掉
type UserInfoResonse struct {
	UserId int64 `json:"user_id"`
}

type UserController struct{}

// 用户信息
func (c *UserController) UserInfo(ctx *gin.Context) {
	var req requests.UserInfoRequest
	if ok := requests.Validate(ctx, &req, requests.UserInfoValid); !ok {
		return
	}
	userService := app.UserService()
	userInfo, err := userService.GetUserInfo(ctx, 1)
	response.WriteResponse(ctx, err, userInfo)
}
