//Package serializer ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 22:20:33
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 10:58:12
 */
package serializer

import "cmall/model"

// Category 分类序列化器
type Category struct {
	ID           uint   `json:"id"`
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
	CreatedAt    int64  `json:"created_at"`
}

// BuildCategory 序列化分类
func BuildCategory(item model.Category) Category {
	return Category{
		ID:           item.ID,
		CategoryID:   item.CategoryID,
		CategoryName: item.CategoryName,
		CreatedAt:    item.CreatedAt.Unix(),
	}
}

// BuildCategories 序列化分类列表
func BuildCategories(items []model.Category) (categories []Category) {
	for _, item := range items {
		category := BuildCategory(item)
		categories = append(categories, category)
	}
	return categories
}
