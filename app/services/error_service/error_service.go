/**
* @Author: cr-mao
* @Desc: error 测试service
**/
package error_service

import (
	"goframe/app/errcode"
	"goframe/infra/errors"
)

type ErrorService struct{}

func NewErrorService() *ErrorService {
	return &ErrorService{}
}

var errDemo1 = errors.New("demo error")

func (s *ErrorService) Demo1() error {
	var err = errDemo1 // 假设是model层或者其他service 返回的error
	return errors.WrapC(err, errcode.ErrUserInvalid, "ErrorService demo1 get error")
}
