package serializer

import (
	"cmall/model"
)

// Favorites 视频序列化器
type Favorites struct {
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"id"`
	CreatedAt     int64  `json:"created_at"`
	Name          string `json:"name"`
	CategoryID    int    `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
}

// BuildFavorite 序列化收藏夹
func BuildFavorite(item1 model.Favorites, item2 model.Products) Favorites {
	return Favorites{
		UserID:        item1.UserID,
		ProductID:     item1.ProductID,
		CreatedAt:     item1.CreatedAt.Unix(),
		Name:          item2.Name,
		CategoryID:    item2.CategoryID,
		Title:         item2.Title,
		Info:          item2.Info,
		ImgPath:       item2.ImgPath,
		Price:         item2.Price,
		DiscountPrice: item2.DiscountPrice,
	}
}

// BuildFavorites 序列化收藏夹列表
func BuildFavorites(items []model.Favorites) (favorites []Favorites) {
	for _, item1 := range items {
		item2 := model.Products{}
		err := model.DB.First(&item2, item1.ProductID).Error
		if err != nil {
			continue
		}
		favorite := BuildFavorite(item1, item2)
		favorites = append(favorites, favorite)
	}
	return favorites
}
