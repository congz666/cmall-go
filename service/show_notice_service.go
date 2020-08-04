//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 23:12:42
 * @LastEditors: congz
 * @LastEditTime: 2020-08-04 11:04:02
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// ShowNoticeService 公告详情服务
type ShowNoticeService struct {
}

// Show 公告详情服务
func (service *ShowNoticeService) Show() serializer.Response {
	var notice model.Notice
	code := e.SUCCESS
	if err := model.DB.First(&notice).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildNotice(notice),
	}
}
