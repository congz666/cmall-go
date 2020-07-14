package serializer

import "cmall/model"

// Products 视频序列化器
type Products struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryID    int    `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint64 `json:"view"`
	CreatedAt     int64  `json:"created_at"`
}

// BuildProduct 序列化商品
func BuildProduct(item model.Products) Products {
	return Products{
		ID:            item.ID,
		Name:          item.Name,
		CategoryID:    item.CategoryID,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		View:          item.View(),
		CreatedAt:     item.CreatedAt.Unix(),
	}
}

// BuildProducts 序列化商品列表
func BuildProducts(items []model.Products) (products []Products) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}
	return products
}
