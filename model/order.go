//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 11:46:28
 * @LastEditors: congz
 * @LastEditTime: 2020-08-05 14:36:17
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// Order 订单模型
type Order struct {
	gorm.Model
	UserID       uint
	ProductID    uint
	Num          uint
	OrderNum     uint64
	AddressName  string
	AddressPhone string
	Address      string
	Status       uint
}
