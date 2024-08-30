/**
* @Author: cr-mao
* @Desc:  用户任务服务
**/
package user_task_service

import "goframe/app/services/user_service"

type UserTaskService struct {
	UserService *user_service.UserService
}

func NewUserTaskService(userService *user_service.UserService) *UserTaskService {
	return &UserTaskService{
		UserService: userService,
	}
}
