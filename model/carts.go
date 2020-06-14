package model

import (
	"github.com/jinzhu/gorm"
)

// Carts 订单模型
type Carts struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Num       uint
	MaxNum    uint
	Check     bool
}
