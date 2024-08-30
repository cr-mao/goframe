## http错误码设计


	项目组代号:10
	服务代号:01
	模块代号:0~99
	错误码：0~99
	| 错误标识                | 错误码   | HTTP状态码 | 描述                          |
	| ----------------------- | -------- | ---------- | ----------------------------- |
	| ErrNo                   |  0 | 200        |  OK                            |
	| ErrInternalServer       | 10010001 | 500        |  Internal server error （服务器内部错误）      |
	| ErrParams               | 10010002 | 400        |  Illegal params  (请求参数不合法)                |
	| ErrAuthenticationHeader | 10010003 | 401        |  Authentication header Illegal  (要登录的接口，头的token认证失败,失败跳登录页面)|
	| ErrAuthentication       | 10010004 | 401        |  Authentication failed  (登录失败，输入账户、密码失败)|
	| ErrNotFound             | 10010005 | 404        |  Route not found     (请求路由找不到）             |
	| ErrPermission           | 10010006 | 403        |  Permission denied (没有权限,一些接口可能没请求权限)            |
	| ErrTooFast              | 10010007 | 429        |  Too Many Requests （用户在给定的时间内发送了太多请求）            |
	| ErrTimeout              | 10010008 | 504        |  Server response timeout   （go服务这边不会返回，一般是nginx、网关超时 才返回504）|
	| ErrMysqlServer          | 10010101 | 500        |  Mysql server error      （mysql 服务错误)       |
	| ErrMysqlSQL             | 10010102 | 500        |  Illegal SQL               (sql 代码错误）       |
	| ErrRedisServer          | 10010201 | 500        |  Redis server error        （redis 服务错误）    |



使用示例
目前错误码是和error在一起的,返回错误的使用，如果是业务错误，请在app/errcode定义错误,

框架基础错误见 app/errcode/base_module.go

```go
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
```


成功响应,http_status_code=200
```json
{
    "error_code":0,
    "msg":"",
    "data": {
      "user_name":"mzy"
    }
}
```

错误响应,如http_status_code=500
```json
{
    "error_code": 10010001,
    "msg":"Internal server error",
    "data": null
}
```



