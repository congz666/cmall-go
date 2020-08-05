//Package serializer ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 13:39:25
 * @LastEditors: congz
 * @LastEditTime: 2020-08-05 14:45:18
 */
package serializer

import (
	"cmall/model"
)

// Order 订单序列化器
type Order struct {
	ID            uint   `json:"id"`
	OrderNum      uint64 `json:"order_num"`
	CreatedAt     int64  `json:"created_at"`
	UserID        uint   `json:"user_id"`
	Num           uint   `json:"num"`
	AddressName   string `json:"address_name"`
	AddressPhone  string `json:"address_phone"`
	Address       string `json:"address"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
	DiscountPrice string `json:"discount_price"`
}

// BuildOrder 序列化收藏夹
func BuildOrder(item1 model.Order, item2 model.Product) Order {
	return Order{
		ID:            item1.ID,
		OrderNum:      item1.OrderNum,
		CreatedAt:     item1.CreatedAt.Unix(),
		UserID:        item1.UserID,
		Num:           item1.Num,
		AddressName:   item1.AddressName,
		AddressPhone:  item1.AddressPhone,
		Address:       item1.Address,
		Name:          item2.Name,
		ImgPath:       item2.ImgPath,
		DiscountPrice: item2.DiscountPrice,
	}
}

// BuildOrders 序列化收藏夹列表
func BuildOrders(items []model.Order) (orders []Order) {
	for _, item1 := range items {
		item2 := model.Product{}
		err := model.DB.First(&item2, item1.ProductID).Error
		if err != nil {
			continue
		}
		order := BuildOrder(item1, item2)
		orders = append(orders, order)
	}
	return orders
}
