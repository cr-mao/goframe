// Package requests 处理请求数据和表单验证
package requests

import (
	"goframe/app/errcode"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"

	"goframe/app/http/response"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

// Validate 控制器里调用示例：
//
//	if ok := requests.Validate(ctx, &requests.DemoRequest{}, requests.DemoValidFunc); !ok {
//		return
//	}
func Validate(ctx *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	//  1. 解析请求，ShouldBind 支持 JSON 数据、表单请求和 URL Query, ShouldBindJSON 支持json
	if err := ctx.ShouldBindJSON(obj); err != nil {
		response.Write(ctx, 400, response.Response{
			ErrorCode: errcode.ErrParams,
			Msg:       "Request parse error ,body should be json and type should be consistent",
			Data:      nil,
		})
		return false
	}
	// 2. 表单验证
	errs := handler(obj, ctx)
	// 3. 判断验证是否通过
	if len(errs) > 0 {
		errorStr := ""
		//map[string][]string
		for k, v := range errs {
			errorStr += k + ":" + strings.Join(v, " , ") + ";"
		}
		errorStr = strings.Trim(errorStr, ";")
		response.Write(ctx, 400, response.Response{
			ErrorCode: errcode.ErrParams,
			Msg:       errorStr,
			Data:      nil,
		})
		return false
	}
	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}
	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
