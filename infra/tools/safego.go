/**
* @Author: cr-mao
* @Desc: 安全开启go程
**/
package tools

import (
	"go.uber.org/zap"

	"goframe/infra/logger"
)

// 安全开启野生goroutine, go SafeGo(fn)
func SafeGo(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("run go err", zap.Any("err", err))
		}
	}()
	fn()
}
