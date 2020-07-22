//Package serializer ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 13:39:25
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 14:54:29
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
	ProductID     uint   `json:"product_id"`
	Num           uint   `json:"num"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
	DiscountPrice string `json:"discount_price"`
}

// OrderDetails 订单详情序列化器
type OrderDetails struct {
	Order   interface{} `json:"order"`
	Address interface{} `json:"address"`
}

// BuildOrder 序列化收藏夹
func BuildOrder(item1 model.Order, item2 model.Product) Order {
	return Order{
		ID:            item1.ID,
		OrderNum:      item1.OrderNum,
		CreatedAt:     item1.CreatedAt.Unix(),
		UserID:        item1.UserID,
		ProductID:     item1.ProductID,
		Num:           item1.Num,
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
