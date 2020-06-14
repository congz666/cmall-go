package serializer

import (
	"cmall/model"
)

// Carts 购物车序列化器
type Carts struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	Num           uint   `json:"num"`
	MaxNum        uint   `json:"max_num"`
	Check         bool   `json:"check"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
	DiscountPrice string `json:"discount_price"`
}

// BuildCart 序列化购物车
func BuildCart(item1 model.Carts, item2 model.Products) Carts {
	return Carts{
		ID:            item1.ID,
		UserID:        item1.UserID,
		ProductID:     item1.ProductID,
		CreatedAt:     item1.CreatedAt.Unix(),
		Num:           item1.Num,
		MaxNum:        item1.MaxNum,
		Check:         item1.Check,
		Name:          item2.Name,
		ImgPath:       item2.ImgPath,
		DiscountPrice: item2.DiscountPrice,
	}
}

// BuildCarts 序列化购物车列表
func BuildCarts(items []model.Carts) (carts []Carts) {
	for _, item1 := range items {
		item2 := model.Products{}
		err := model.DB.First(&item2, item1.ProductID).Error
		if err != nil {
			continue
		}
		cart := BuildCart(item1, item2)
		carts = append(carts, cart)
	}
	return carts
}
