//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 20:04:25
 * @LastEditors: congz
 * @LastEditTime: 2020-08-04 11:25:36
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// UpdateNoticeService 公告更新的服务
type UpdateNoticeService struct {
	Text string `form:"text" json:"text"`
}

// Update 公告更新的服务
func (service *UpdateNoticeService) Update() serializer.Response {
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
	notice.Text = service.Text
	if err := model.DB.Save(&notice).Error; err != nil {
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
	}
}
