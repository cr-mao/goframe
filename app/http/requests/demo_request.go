/**
* @Author: cr-mao
* @Desc: demo 请求验证
**/
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type UserInfoRequest struct {
	UserId   int64  `valid:"user_id" json:"user_id"`
	UserName string `valid:"user_name" json:"user_name"`
}

func UserInfoValid(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"user_id":   []string{"required"},
		"user_name": []string{"required", "between:3,8"},
	}
	messages := govalidator.MapData{
		"user_id": []string{
			"required:must be not empty",
		},
		"user_name": []string{
			"required:must be not empty",
			"between: 用户名长度必须为3到8个字符",
		},
	}
	return validate(data, rules, messages)
}
