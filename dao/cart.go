package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func (dao *CartDao) CreateCart(in *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Create(&in).Error
}

func (dao *CartDao) UpdateCartNum(uid, cid uint, num int) error {
	return dao.DB.Model(&model.Cart{}).Where("id = ? AND user_id = ?", cid, uid).Update("num", num).Error
}

func (dao *CartDao) ListCartByUserId(uid uint) (cartes []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id = ?", uid).Find(&cartes).Error
	return
}

func (dao *CartDao) DeleteCartByCid(uid, cid uint) error {
	return dao.DB.Model(&model.Cart{}).Where("id = ? AND user_id = ?", cid, uid).Delete(&model.Cart{}).Error
}
