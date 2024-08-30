//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"goframe/app/services/error_service"

	"goframe/app/services/user_service"
	"goframe/app/services/user_task_service"
)

// 请按字母排序
var SuperSet = wire.NewSet(
	// e
	error_service.NewErrorService,
	// u
	user_service.NewUserService,
	user_task_service.NewUserTaskService,
)

// 用户服务
func UserService() *user_service.UserService {
	wire.Build(SuperSet)
	return &user_service.UserService{}
}

// 用户任务服务
func UserTaskService() *user_task_service.UserTaskService {
	wire.Build(SuperSet)
	return &user_task_service.UserTaskService{}
}

func ErrorService() *error_service.ErrorService {
	wire.Build(SuperSet)
	return &error_service.ErrorService{}
}
