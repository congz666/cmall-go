package serializer

import "cmall/model"

// Carousels 轮播图序列化器
type Carousels struct {
	ID        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	CreatedAt int64  `json:"created_at"`
}

// BuildCarousel 序列化轮播图
func BuildCarousel(item model.Carousels) Carousels {
	return Carousels{
		ID:        item.ID,
		ImgPath:   item.ImgPath,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildCarousels 序列化轮播图列表
func BuildCarousels(items []model.Carousels) (carousels []Carousels) {
	for _, item := range items {
		carousel := BuildCarousel(item)
		carousels = append(carousels, carousel)
	}
	return carousels
}
