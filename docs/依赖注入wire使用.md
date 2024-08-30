## 依赖注入wire工具


download

```shell
go install github.com/google/wire/cmd/wire@v0.6.0
```

在app/wire.go 新增依赖service， 这里如user_service 依赖user_task_service .

```go
/**
* @Author: cr-mao
* @Desc: 用户service
**/
package user_service

import "goframe/app/services/user_task_service"

type UserService struct {
	UserTaskService *user_task_service.UserTaskService
}

func NewUserService(userTaskService *user_task_service.UserTaskService) *UserService {
	return &UserService{
		UserTaskService: userTaskService,
	}
}
```

编辑wire.go 
```go
//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"

	"goframe/app/services/user_service"
	"goframe/app/services/user_task_service"
)

// 请按字母排序
var SuperSet = wire.NewSet(
	// u
	user_service.NewUserService,
	user_task_service.NewUserTaskService,
)

func UserService() *user_service.UserService {
	wire.Build(SuperSet)
	return &user_service.UserService{}
}

func UserTaskService() *user_task_service.UserTaskService {
	wire.Build(SuperSet)
	return &user_task_service.UserTaskService{}
}
```


生成依赖
```shell
make wire 
```


使用
```text
userService :=app.UserService() 
userService.UserTaskService.Doxxx()
```



