//Package serializer ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 15:33:11
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 10:57:43
 */
package serializer

import (
	"cmall/model"
)

// Cart 购物车序列化器
type Cart struct {
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
func BuildCart(item1 model.Cart, item2 model.Product) Cart {
	return Cart{
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
func BuildCarts(items []model.Cart) (carts []Cart) {
	for _, item1 := range items {
		item2 := model.Product{}
		err := model.DB.First(&item2, item1.ProductID).Error
		if err != nil {
			continue
		}
		cart := BuildCart(item1, item2)
		carts = append(carts, cart)
	}
	return carts
}
