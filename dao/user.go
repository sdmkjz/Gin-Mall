package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExistOrNotByUserName 根据username 判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(username string) (user *model.User, exist bool, err error) {
	var count int64
	if err = dao.DB.Model(&model.User{}).Where("user_name=?", username).Find(&user).Count(&count).Error; count == 0 {
		return nil, false, err
	}
	return user, true, nil
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	return dao.DB.Model(&user).Create(&user).Error
}

// GetUserById 根据ID获取User
func (dao *UserDao) GetUserById(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

// UpdateUserById 通过ID更新user信息
func (dao *UserDao) UpdateUserById(uid uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id = ?", uid).Updates(&user).Error
}
