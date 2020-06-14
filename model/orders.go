package model

import (
	"github.com/jinzhu/gorm"
)

// Orders 订单模型
type Orders struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Num       uint
	OrderID   uint64
}
