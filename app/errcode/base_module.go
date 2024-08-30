/**
* @Author: cr-mao
* @Desc: 基础模块相关errcode
**/
package errcode

/*
*
	项目组代号:10
	服务代号:01
	模块代号:0~99
	错误码：0~99
	| 错误标识                | 错误码   | HTTP状态码 | 描述                          |
	| ----------------------- | -------- | ---------- | ----------------------------- |
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

客户端要关心的是http code 是     ErrParams  400 ，ErrAuthenticationHeader 401，ErrAuthentication 401，ErrPermission 403, ErrTooFast 429
当返回http code= 401 时，   ErrAuthenticationHeader,ErrAuthentication 根据具体Code 去区分处理
*/

const (
	ErrInternalServer       = 10010001 // Internal server error （服务器内部错误）
	ErrParams               = 10010002 // Illegal params  (请求参数不合法)
	ErrAuthenticationHeader = 10010003 // Authentication header Illegal  (要登录)
	ErrAuthentication       = 10010004 // 登录失败，输入账户、密码失败
	ErrRouteNotFound        = 10010005 // (请求路由找不到）
	ErrPermission           = 10010006 // 没有权限
	ErrTooFast              = 10010007 // 请求太快
	ErrMysqlServer          = 10010008 // Mysql server error
	ErrRedisServer          = 10010009 // Redis server error
)

func registerBaseCode() {
	register(ErrInternalServer, 500, "Internal server error")
	register(ErrParams, 400, "Illegal params")
	register(ErrAuthenticationHeader, 401, "Authentication header Illegal")
	register(ErrAuthentication, 401, "Authentication failed")
	register(ErrRouteNotFound, 404, "Route not found")
	register(ErrPermission, 403, "Permission denied")
	register(ErrTooFast, 429, "Too Many Requests")
	register(ErrMysqlServer, 500, "Mysql server error")
	register(ErrRedisServer, 500, "redis server error")
}
