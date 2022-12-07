package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCategoryByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

// 根据ID获取Category
func (dao *CarouselDao) ListCategory() (category []model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&category).Error
	return category, err
}
