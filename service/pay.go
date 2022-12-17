package service

import (
	"context"
	"errors"
	"fmt"
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/pkg/utils"
	"gin-mall/serializer"
	"strconv"
)

type OrderPay struct {
	OrderId   uint   `json:"order_id" form:"order_id"`
	Money     int64  `json:"money" form:"money"`
	OrderNo   string `json:"order_no" form:"order_no"`
	ProductId uint   `json:"product_id" form:"product_id"`
	PayTime   string `json:"pay_time" form:"pay_time"`
	Sign      string `json:"sign" form:"sign"`
	BossId    uint   `json:"boss_id" form:"boss_id"`
	BossName  string `json:"boss_name" form:"boss_name"`
	Num       int    `json:"num" form:"num"`
	Key       string `json:"key" form:"key"`
}

func (service *OrderPay) PayDown(ctx context.Context, uid uint) serializer.Response {
	utils.Encrypt.SetKey(service.Key)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	tx := orderDao.Begin()
	order, err := orderDao.GetOrderByOid(uid, service.OrderId)
	if err != nil {
		utils.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	money := order.Money
	num := order.Num
	money = money * float64(num)

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uid)
	if err != nil {
		utils.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 对金额进行解密 减去订单 再加密保存
	moneyStr := utils.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)
	if moneyFloat-money < 0.0 {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("金额不足").Error(),
		}
	}
	finMoney := fmt.Sprintf("%f", moneyFloat-money)
	user.Money = utils.Encrypt.AesEncoding(finMoney)

	userDao = dao.NewUserDaoByDB(userDao.DB)
	err = userDao.UpdateUserById(uid, user)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("金额不足").Error(),
		}
	}
	// 商家收入
	var boss *model.User
	boss, err = userDao.GetUserById(service.BossId)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	moneyStr = utils.Encrypt.AesDecoding(boss.Money)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	finMoney = fmt.Sprintf("%f", moneyFloat+money)
	boss.Money = utils.Encrypt.AesEncoding(finMoney)

	err = userDao.UpdateUserById(boss.ID, boss)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 对应商品数量 -
	var product *model.Product
	productDao := dao.NewProductDao(ctx)
	product, err = productDao.GetProductById(service.ProductId)
	product.Num -= num
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = productDao.UpdateProduct(service.ProductId, product)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 订单删除
	err = orderDao.DeleteOrderByOrderId(uid, service.OrderId)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 自己的商品+1
	productUser := &model.Product{
		Name:          product.Name,
		CategoryId:    product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.Price,
		OnSale:        false,
		Num:           1,
		BossId:        uid,
		BossName:      user.UserName,
		BossAvatar:    user.Avatar,
	}
	err = productDao.CreateProduct(productUser)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	tx.Commit()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
