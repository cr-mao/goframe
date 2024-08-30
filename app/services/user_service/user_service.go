/**
* @Author: cr-mao
* @Desc: 用户service
**/
package user_service

import (
	"context"
	"goframe/app/models/user_model"

	"goframe/app/errcode"
	"goframe/app/services/user_service/dto"
	"goframe/global"
	"goframe/infra/db"
	"goframe/infra/errors"
)

// 用户service,依赖usertask服务, 这里写错了。 其实是任务服务依赖用户服务。 只是个demo。
type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// GetUserInfo
//
//	@Description:  用户信息
//	@receiver s
//	@param ctx
//	@param userId
//	@return *dto.UserDto  返回数据
//	@return error
func (s *UserService) GetUserInfo(ctx context.Context, userId int64) (*dto.UserDto, error) {
	session := db.Client(ctx, global.DB_DEFAULT)
	userRow := user_model.GetByUserId(session, userId)
	if userRow.UserId == 0 {
		return nil, errors.WithCode(errcode.ErrUserNotFound, "user not found %d", userId)
	}
	var res dto.UserDto
	res.UserId = userRow.UserId
	res.Guid = userRow.Guid
	return &res, nil
}
