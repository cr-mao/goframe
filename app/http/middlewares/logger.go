// 请求错误log记录
package middlewares

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"

	"goframe/infra/helpers"
	"goframe/infra/logger"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// Logger 记录请求日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 response 内容
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w
		// 获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body 是一个 buffer 对象，只能读取一次
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 读取后，重新赋值 c.Request.Body ，以供后续的其他操作
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 设置开始时间
		start := time.Now()
		c.Next()
		responStatus := c.Writer.Status()
		if responStatus >= 400 {
			// 开始记录日志的逻辑
			cost := time.Since(start)
			costTime := helpers.GetMicroseconds(cost)
			responseContent := w.body.String()
			requestBodyContent := string(requestBody)

			logFields := []zap.Field{
				zap.Int("status", responStatus),
				zap.String("request_uri", c.Request.RequestURI),
				//zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Int64("cost_time", costTime),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
			}
			errDetail, ok := c.Get("error_detail")
			if ok {
				logFields = append(logFields, zap.String("errors", errDetail.(string)))
			}
			logFields = append(logFields, zap.String("body", requestBodyContent), zap.String("response", responseContent))
			logger.Error("HTTP Error "+cast.ToString(responStatus), logFields...)
		}
	}
}
