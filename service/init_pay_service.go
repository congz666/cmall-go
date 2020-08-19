//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 10:58:11
 * @LastEditors: congz
 * @LastEditTime: 2020-08-17 11:35:53
 */
package service

import (
	"bytes"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// InitPayService 初始化支付的服务
type InitPayService struct {
	OrderNum string `form:"order_num" json:"order_num" `
	PayType  string `form:"pay_type" json:"pay_type" `
	Amount   string `form:"amount" json:"amount"`
}

// PayOrderInfo PayOrderInfo
type PayOrderInfo struct {
	ID     string `json:"id"`     //渠道唯一ID
	PayURL string `json:"payUrl"` //支付页URL
}

// Result ...
type Result struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data PayOrderInfo `json:"data"`
}

// Init 初始化支付 详情请查阅FM支付文档
func (service *InitPayService) Init() serializer.Response {
	code := e.SUCCESS
	//计算签名
	var buff bytes.Buffer
	buff.WriteString(os.Getenv("FM_Pay_ID"))
	buff.WriteString(service.OrderNum)
	buff.WriteString(service.Amount)
	buff.WriteString(os.Getenv("FM_Pay_NotifyURL"))
	buff.WriteString(os.Getenv("FM_Pay_Key"))
	sign := fmt.Sprintf("%x", md5.Sum(buff.Bytes()))

	returnURL := os.Getenv("FM_Pay_ReturnURL") + service.OrderNum
	//构造请求参数
	buff.Reset()
	buff.WriteString("sign=")
	buff.WriteString(sign)
	buff.WriteString("&amount=")
	buff.WriteString(service.Amount)
	buff.WriteString("&orderNo=")
	buff.WriteString(service.OrderNum)
	buff.WriteString("&payType=")
	buff.WriteString(service.PayType)
	buff.WriteString("&merchantNum=")
	buff.WriteString(os.Getenv("FM_Pay_ID"))
	buff.WriteString("&notifyUrl=")
	buff.WriteString(os.Getenv("FM_Pay_NotifyURL"))
	buff.WriteString("&returnUrl=")
	buff.WriteString(returnURL)
	buff.WriteString("&attch=")
	buff.WriteString(os.Getenv("FM_Pay_attch"))
	//调用渠道接口
	resp, err := http.Post("http://zfapi.nnt.ltd/api/startOrder", "application/x-www-form-urlencoded", &buff)
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
	var result Result
	if err = json.Unmarshal(data, &result); err != nil {
		logging.Info(err)
		code = e.ERROR_UNMARSHAL_JSON
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: result.Code,
		Msg:    result.Msg,
		Data:   result.Data,
	}
}
