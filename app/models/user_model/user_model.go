// Package user 模型
package user_model

type User struct {
	//fix
	UserId int64 `gorm:"column:user_id;primaryKey;autoIncrement;" json:"user_id,omitempty"`
	// user_id
	// 用户唯一标志 即token
	Guid string `gorm:"column:guid;index;" json:"guid"`
	// 0正常1禁用
	ForbiddenStatus int64 `gorm:"column:forbidden_status" json:"forbidden_status"`
}

func (User) TableName() string {
	return "user"
}
