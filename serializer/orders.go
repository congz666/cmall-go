package serializer

import (
	"cmall/model"
)

// Orders 视频序列化器
type Orders struct {
	ID            uint   `json:"id"`
	OrderID       uint64 `json:"order_id"`
	CreatedAt     int64  `json:"created_at"`
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	Num           uint   `json:"num"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
	DiscountPrice string `json:"discount_price"`
}

// BuildOrder 序列化收藏夹
func BuildOrder(item1 model.Orders, item2 model.Products) Orders {
	return Orders{
		ID:            item1.ID,
		OrderID:       item1.OrderID,
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
func BuildOrders(items []model.Orders) (orders []Orders) {
	for _, item1 := range items {
		item2 := model.Products{}
		err := model.DB.First(&item2, item1.ProductID).Error
		if err != nil {
			continue
		}
		order := BuildOrder(item1, item2)
		orders = append(orders, order)
	}
	return orders
}
