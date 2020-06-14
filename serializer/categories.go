package serializer

import "cmall/model"

// Categories 分类序列化器
type Categories struct {
	ID           uint   `json:"id"`
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
	CreatedAt    int64  `json:"created_at"`
}

// BuildCategory 序列化分类
func BuildCategory(item model.Categories) Categories {
	return Categories{
		ID:           item.ID,
		CategoryID:   item.CategoryID,
		CategoryName: item.CategoryName,
		CreatedAt:    item.CreatedAt.Unix(),
	}
}

// BuildCategories 序列化分类列表
func BuildCategories(items []model.Categories) (categories []Categories) {
	for _, item := range items {
		category := BuildCategory(item)
		categories = append(categories, category)
	}
	return categories
}
