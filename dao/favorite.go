package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteDaoByDb(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}

func (dao *FavoriteDao) ListFavorites(uid uint) (resp []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id = ?", uid).Find(&resp).Error
	return
}

func (dao *FavoriteDao) FavoriteExistOrNot(pid, uid uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("product_id = ? AND user_id = ?", pid, uid).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, err
	}
	return true, nil
}

func (dao *FavoriteDao) CreateFavorite(in *model.Favorite) error {
	return dao.DB.Model(&model.Favorite{}).Create(&in).Error
}

func (dao *FavoriteDao) DeleteFavorite(uid, fid uint) error {
	return dao.DB.Model(&model.Favorite{}).Where("id = ? AND user_id = ?", fid, uid).Delete(&model.Favorite{}).Error
}
