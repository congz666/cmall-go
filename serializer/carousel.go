//Package serializer ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 14:14:38
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 10:57:06
 */
package serializer

import "cmall/model"

// Carousel 轮播图序列化器
type Carousel struct {
	ID        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	CreatedAt int64  `json:"created_at"`
}

// BuildCarousel 序列化轮播图
func BuildCarousel(item model.Carousel) Carousel {
	return Carousel{
		ID:        item.ID,
		ImgPath:   item.ImgPath,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildCarousels 序列化轮播图列表
func BuildCarousels(items []model.Carousel) (carousels []Carousel) {
	for _, item := range items {
		carousel := BuildCarousel(item)
		carousels = append(carousels, carousel)
	}
	return carousels
}
