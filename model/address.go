package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	Name    string `type:"varchar(20) not null"`
	Phone   string `type:"varchar(11) not null"`
	Address string `type:"varchar(50) not null"`
}
