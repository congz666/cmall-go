//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-20 14:40:45
 * @LastEditors: congz
 * @LastEditTime: 2020-08-19 11:04:35
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/pkg/util"
	"cmall/serializer"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// VaildQQService 查询，一键注册，绑定服务
type VaildQQService struct {
	AuthorizationCode string `form:"authorization_code" json:"authorization_code"`
}

// QQTokenResult 调用QQAPI返回的数据
type QQTokenResult struct {
	Code         int    `json:"code"`
	Msg          string `json:"msg"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// QQIDResult 调用QQAPI返回的数据
type QQIDResult struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	ClientID string `json:"client_id"`
	OpenID   string `json:"openid"`
}

// Vaild 绑定邮箱
func (service *VaildQQService) Vaild() serializer.Response {
	code := e.SUCCESS
	grantType := "authorization_code"
	fmtType := "json"
	path := os.Getenv("QQ_Redirect_URI")
	redirectURI := url.QueryEscape(path)
	//获取AccessToken
	getTokenURL := fmt.Sprintf("https://graph.qq.com/oauth2.0/token?grant_type=%s&client_id=%s&client_secret=%s&code=%s&redirect_uri=%s&fmt=%s",
		grantType, os.Getenv("QQ_Client_ID"), os.Getenv("QQ_Client_KEY"), service.AuthorizationCode, redirectURI, fmtType)
	resp, err := http.Get(getTokenURL)
	if err != nil {
		logging.Info(err)
		code = e.ERROR_CALL_API
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//读取所有响应数据
	var data []byte
	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		logging.Info(err)
		code = e.ERROR_READ_FILE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//解析渠道返回的JSON
	var result QQTokenResult
	if err = json.Unmarshal(data, &result); err != nil {
		logging.Info(err)
		code = e.ERROR_UNMARSHAL_JSON
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//获取OpenID
	getIDURL := fmt.Sprintf("https://graph.qq.com/oauth2.0/me?access_token=%s&fmt=%s",
		result.AccessToken, fmtType)
	resp, err = http.Get(getIDURL)
	if err != nil {
		logging.Info(err)
		code = e.ERROR_CALL_API
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//读取所有响应数据
	var data2 []byte
	if data2, err = ioutil.ReadAll(resp.Body); err != nil {
		logging.Info(err)
		code = e.ERROR_READ_FILE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//解析渠道返回的JSON
	var result2 QQIDResult
	if err = json.Unmarshal(data2, &result2); err != nil {
		logging.Info(err)
		code = e.ERROR_UNMARSHAL_JSON
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	var user model.User
	if err := model.DB.Where("id=?", model.DB.Table("user_auth").Select("user_id").Where("identifier=?", result2.OpenID).SubQuery()).First(&user).Error; err != nil {
		//如果查询不到，则创建CMall账号
		if gorm.IsRecordNotFoundError(err) {
			if result.AccessToken == "" {
				code = e.ERROR_AUTH
				return serializer.Response{
					Status: code,
					Msg:    e.GetMsg(code),
					Error:  err.Error(),
				}
			}
			user.UserName = strconv.FormatInt(time.Now().Unix(), 10)
			user.Nickname = user.UserName
			user.Status = model.Active
			user.Avatar = "img/avatar/avatar1.jpg"
			// 创建用户
			if err := model.DB.Create(&user).Error; err != nil {
				logging.Info(err)
				code = e.ERROR_DATABASE
				return serializer.Response{
					Status: code,
					Msg:    e.GetMsg(code),
				}
			}
			//创建第三方登录信息
			userAuth := model.UserAuth{
				UserID:       user.ID,
				IdentityType: "qq",
				Identifier:   result2.OpenID,
				Token:        result.AccessToken,
				RefreshToken: result.RefreshToken,
			}

			if err := model.DB.Create(&userAuth).Error; err != nil {
				logging.Info(err)
				code = e.ERROR_DATABASE
				return serializer.Response{
					Status: code,
					Msg:    e.GetMsg(code),
				}
			}
			token, err := util.GenerateToken(user.UserName, user.PasswordDigest, 0)
			if err != nil {
				logging.Info(err)
				code = e.ERROR_AUTH_TOKEN
				return serializer.Response{
					Status: code,
					Msg:    e.GetMsg(code),
				}
			}
			return serializer.Response{
				Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//获取token
	token, err := util.GenerateToken(user.UserName, user.PasswordDigest, 0)
	if err != nil {
		logging.Info(err)
		code = e.ERROR_AUTH_TOKEN
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
