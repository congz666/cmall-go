package service

import (
	"cmall/model"
	"cmall/serializer"
	"strconv"
)

// ShowCategoryService 视频列表服务
type ShowCategoryService struct {
}

// Show 分类列表
func (service *ShowCategoryService) Show(CategoryID string) serializer.Response {
	products := []model.Products{}
	id, err := strconv.Atoi(CategoryID)
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "字符串转换错误",
			Error:  err.Error(),
		}
	}
	if err := model.DB.Where("category_id=?", id).Find(&products).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildProducts(products),
	}
}
