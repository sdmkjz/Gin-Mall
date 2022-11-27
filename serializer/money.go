package serializer

import (
	"gin-mall/model"
	"gin-mall/pkg/utils"
)

type Money struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(item *model.User, key string) Money {
	utils.Encrypt.SetKey(key)
	return Money{
		UserId:    item.ID,
		UserName:  item.UserName,
		UserMoney: utils.Encrypt.AesDecoding(item.Money),
	}
}
