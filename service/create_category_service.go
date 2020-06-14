package service

import (
	"cmall/model"
	"cmall/serializer"
)

// CreateCategoryService 收藏创建的服务
type CreateCategoryService struct {
	CategoryID   uint   `form:"category_id" json:"category_id"`
	CategoryName string `form:"category_name" json:"category_name"`
}

// Create 创建分类
func (service *CreateCategoryService) Create() serializer.Response {
	category := model.Categories{
		CategoryID:   service.CategoryID,
		CategoryName: service.CategoryName,
	}

	err := model.DB.Create(&category).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "创建分类失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildCategory(category),
	}
}
