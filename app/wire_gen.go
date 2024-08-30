// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/google/wire"
	"goframe/app/services/error_service"
	"goframe/app/services/user_service"
	"goframe/app/services/user_task_service"
)

// Injectors from wire.go:

// 用户服务
func UserService() *user_service.UserService {
	userService := user_service.NewUserService()
	return userService
}

// 用户任务服务
func UserTaskService() *user_task_service.UserTaskService {
	userService := user_service.NewUserService()
	userTaskService := user_task_service.NewUserTaskService(userService)
	return userTaskService
}

func ErrorService() *error_service.ErrorService {
	errorService := error_service.NewErrorService()
	return errorService
}

// wire.go:

// 请按字母排序
var SuperSet = wire.NewSet(error_service.NewErrorService, user_service.NewUserService, user_task_service.NewUserTaskService)
