package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func (dao *OrderDao) CreateOrder(in *model.Order) error {
	return dao.DB.Model(&model.Order{}).Create(&in).Error
}

func (dao *OrderDao) GetOrderByOid(uid uint, oid uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id = ? AND user_id = ?", oid, uid).First(&order).Error
	return
}

func (dao *OrderDao) GetOrderByUserId(uid uint) (order []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("user_id = ?", uid).Find(&order).Error
	return
}

func (dao *OrderDao) GetOrderByCondition(condition map[string]interface{}, page model.BasePage) (order []*model.Order, total int64, err error) {
	err = dao.DB.Model(&model.Order{}).
		Where(condition).Count(&total).Error
	if err != nil {
		return
	}
	err = dao.DB.Model(&model.Order{}).
		Where(condition).Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).
		Find(&order).Error
	return
}

func (dao *OrderDao) DeleteOrderByOrderId(uid uint, oid uint) error {
	return dao.DB.Model(&model.Order{}).Where("id = ? AND user_id = ?", oid, uid).Delete(&model.Order{}).Error
}
