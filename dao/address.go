package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func (dao *AddressDao) CreateAddress(in *model.Address) error {
	return dao.DB.Model(&model.Address{}).Create(&in).Error
}

func (dao *AddressDao) GetAddressByAid(aid, uid uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id = ? AND user_id = ?", aid, uid).First(&address).Error
	return
}

func (dao *AddressDao) ListAddressByUserId(uid uint) (address []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("user_id = ?", uid).Find(&address).Error
	return
}

func (dao *AddressDao) UpdateAddressByUserId(aid, uid uint, address *model.Address) error {
	return dao.DB.Model(&model.Address{}).Where("id = ? AND user_id = ?", aid, uid).Updates(&address).Error
}

func (dao *AddressDao) DeleteAddressByAid(aid, uid uint) error {
	return dao.DB.Model(&model.Address{}).Where("id = ? AND user_id = ?", aid, uid).Delete(&model.Address{}).Error
}
