/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/8/30
**/
package user_model

import "gorm.io/gorm"

// 获得一条
func GetByUserId(tx *gorm.DB, userId int64) *User {
	var user User
	tx.Model(User{}).Where("user_id =?", userId).First(&user)
	return &user
}

// 创建
func Create(tx *gorm.DB, tusers *User) error {
	return tx.Create(&tusers).Error
}

// 获得用户列表
func GetUserListByUserIds(tx *gorm.DB, userIds []int64) []User {
	var res []User
	tx.Model(User{}).Where("user_id in ?", userIds).Find(&res)
	return res
}

// 删除
func DeleteByUserId(tx *gorm.DB, userId int64) error {
	return tx.Where("user_id = ?", userId).Delete(&User{}).Error
}
