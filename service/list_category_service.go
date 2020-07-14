package service

import (
	"cmall/model"
	"cmall/serializer"
	"strconv"
)

// ShowCategoryService 视频列表服务
type ShowCategoryService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// Show 分类列表
func (service *ShowCategoryService) Show(CategoryID string) serializer.Response {
	products := []model.Products{}

	total := 0

	if service.Limit == 0 {
		service.Limit = 15
	}

	id, err := strconv.Atoi(CategoryID)
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "字符串转换错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Model(model.Products{}).Where("category_id=?", id).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Where("category_id=?", id).Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}
