package service

import (
	"context"
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/serializer"
	"strconv"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func (service *AddressService) Create(ctx context.Context, uid uint) serializer.Response {
	var address *model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address = &model.Address{
		UserID:  uid,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.CreateAddress(address)
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

func (service *AddressService) Show(ctx context.Context, uid uint, aid string) serializer.Response {
	addressId, _ := strconv.Atoi(aid)
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(uint(addressId), uid)
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
		Data:   serializer.BuildAddress(address),
	}
}

func (service *AddressService) List(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	addressList, err := addressDao.ListAddressByUserId(uid)
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
		Data:   serializer.BuildAddresses(addressList),
	}
}

func (service *AddressService) Update(ctx context.Context, uid uint, aid string) serializer.Response {
	var address *model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address = &model.Address{
		UserID:  uid,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	addressId, _ := strconv.Atoi(aid)
	err := addressDao.UpdateAddressByUserId(uint(addressId), uid, address)
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

func (service *AddressService) Delete(ctx context.Context, uid uint, aid string) serializer.Response {
	addressId, _ := strconv.Atoi(aid)
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	err := addressDao.DeleteAddressByAid(uint(addressId), uid)
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
