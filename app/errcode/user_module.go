/**
* @Author: cr-mao
* @Desc: 用户模块相关errcode
**/
package errcode

// 用户模块
const (
	// 业务错误code
	ErrUserNotFound = 10010101 + iota // 用户找不到
	ErrUserInvalid                    // 用户无效
)

func registerUserErrcode() {
	register(ErrUserNotFound, 400, "用户没找到")
	register(ErrUserInvalid, 400, "用户无效")
}
