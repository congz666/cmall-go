//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 10:47:54
 * @LastEditors: congz
 * @LastEditTime: 2020-08-19 11:31:23
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/pkg/util"
	"cmall/serializer"
	"os"
	"strings"

	"gopkg.in/mail.v2"
)

// SendEmailService 发送邮件的服务
type SendEmailService struct {
	UserID   uint   `form:"user_id" json:"user_id"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	//OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}

// Send 发送邮件
func (service *SendEmailService) Send() serializer.Response {
	code := e.SUCCESS
	var address string
	var notice model.Notice
	token, err := util.GenerateEmailToken(service.UserID, service.OperationType, service.Email, service.Password)
	if err != nil {
		logging.Info(err)
		code = e.ERROR_AUTH_TOKEN
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//数据库里 对应邮件id = operation_type+1
	if err := model.DB.First(&notice, service.OperationType+1).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = os.Getenv("VAILD_EMAIL") + token
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "VaildAddress", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_EMAIL"))
	m.SetHeader("To", service.Email)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")抄送
	m.SetHeader("Subject", "CMall")
	m.SetBody("text/html", mailText)

	d := mail.NewDialer(os.Getenv("SMTP_HOST"), 465, os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASS"))
	d.StartTLSPolicy = mail.MandatoryStartTLS

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		logging.Info(err)
		code = e.ERROR_SEND_EMAIL
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
