package service

import (
	"context"
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/serializer"
	"strconv"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       int  `json:"num" form:"num"`
}

func (service *CartService) Create(ctx context.Context, uid uint) serializer.Response {
	var cart *model.Cart
	code := e.Success
	// 判断商品是否存在
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	cartDao := dao.NewCartDao(ctx)
	cart = &model.Cart{
		UserId:    uid,
		ProductId: service.ProductId,
		BossId:    service.BossId,
		Num:       uint(service.Num),
	}
	err = cartDao.CreateCart(cart)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserById(service.BossId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

func (service *CartService) Update(ctx context.Context, uid uint, cid string) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	cartId, _ := strconv.Atoi(cid)
	err := cartDao.UpdateCartNum(uid, uint(cartId), service.Num)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *CartService) Delete(ctx context.Context, uid uint, cid string) serializer.Response {
	cartId, _ := strconv.Atoi(cid)
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	err := cartDao.DeleteCartByCid(uid, uint(cartId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *CartService) List(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	cart, err := cartDao.ListCartByUserId(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarts(ctx, cart),
	}
}
