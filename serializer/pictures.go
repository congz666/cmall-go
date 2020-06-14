package serializer

import "cmall/model"

// Pictures 商品图片序列化器
type Pictures struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
	CreatedAt int64  `json:"created_at"`
}

// BuildPicture 序列化商品图片
func BuildPicture(item model.Pictures) Pictures {
	return Pictures{
		ID:        item.ID,
		ProductID: item.ProductID,
		ImgPath:   item.ImgPath,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildPictures 序列化轮播图列表
func BuildPictures(items []model.Pictures) (pictures []Pictures) {
	for _, item := range items {
		picture := BuildPicture(item)
		pictures = append(pictures, picture)
	}
	return pictures
}
