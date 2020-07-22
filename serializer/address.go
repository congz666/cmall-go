//Package serializer ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-20 10:52:33
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 10:56:33
 */
package serializer

import (
	"cmall/model"
)

// Address 收货地址序列化器
type Address struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Seen      bool   `json:"seen"`
	CreatedAt int64  `json:"created_at"`
}

// BuildAddress 收货地址购物车
func BuildAddress(item model.Address) Address {
	return Address{
		ID:        item.ID,
		UserID:    item.UserID,
		Name:      item.Name,
		Phone:     item.Phone,
		Address:   item.Address,
		Seen:      false,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildAddresses 序列化收货地址列表
func BuildAddresses(items []model.Address) (addresses []Address) {
	for _, item := range items {
		address := BuildAddress(item)
		addresses = append(addresses, address)
	}
	return addresses
}
