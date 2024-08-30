package errcode

import (
	"goframe/infra/errors"
	"net/http"

	"github.com/novalagung/gubrak"
)

type ErrCode struct {
	//错误码
	C int

	//http的状态码
	HTTP int

	//扩展字段
	Ext string

	//引用文档
	Ref string
}

func (e ErrCode) HTTPStatus() int {
	return e.HTTP
}

func (e ErrCode) String() string {
	return e.Ext
}

func (e ErrCode) Reference() string {
	return e.Ref
}

func (e ErrCode) Code() int {
	if e.C == 0 {
		return http.StatusInternalServerError
	}
	return e.C
}

func register(code int, httpStatus int, message string) {
	found, _ := gubrak.Includes([]int{200, 400, 401, 403, 404, 429, 500}, httpStatus)
	if !found {
		panic("http code not in `200, 400, 401, 403, 404,429, 500`")
	}
	//var ref string
	//if len(refs) > 0 {
	//	ref = refs[0]
	//}
	coder := ErrCode{
		C:    code,
		HTTP: httpStatus,
		Ext:  message,
		Ref:  "",
	}

	errors.MustRegister(coder)
}

var _ errors.Coder = (*ErrCode)(nil)

func RegisterCode() {
	// 注册基础errcode
	registerBaseCode()
	// 注册用户模块相关code
	registerUserErrcode()
}
