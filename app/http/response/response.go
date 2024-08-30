package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"goframe/infra/errors"
)

type Response struct {
	ErrorCode int         `json:"error_code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}

// WriteResponse write an error or the response data into http response body.
// It use errors.ParseCoder to parse any error into errors.Coder
// errors.Coder contains error code, user-safe error message and http status code.
// 业务响应
func WriteResponse(ctx *gin.Context, err error, data interface{}) {
	if err != nil {
		// 错误的堆栈详细信息，
		errStr := fmt.Sprintf("%#+v", err)
		// 通过上下文注入错误堆栈
		ctx.Set("error_detail", errStr)
		coder := errors.ParseCoder(err)
		ctx.JSON(coder.HTTPStatus(), Response{
			ErrorCode: coder.Code(),
			Msg:       coder.String(),
			Data:      nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, Response{
		ErrorCode: 0,
		Msg:       "",
		Data:      data,
	})
}

// 写业务此函数不用，目前仅在request数据验证，及recover中使用,路由没找到使用。
func Write(ctx *gin.Context, httpStatus int, response Response) {
	ctx.JSON(httpStatus, response)
	return
}
