//Package e ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-15 14:53:24
 * @LastEditors: congz
 * @LastEditTime: 2020-08-04 11:42:17
 */
package e

// MsgFlags 状态码map
var MsgFlags = map[int]string{
	SUCCESS:                    "ok",
	ERROR:                      "fail",
	INVALID_PARAMS:             "请求参数错误",
	ERROR_EXIST_NICK:           "已存在该昵称",
	ERROR_EXIST_USER:           "已存在该用户名",
	ERROR_NOT_EXIST_USER:       "该用户不存在",
	ERROR_NOT_COMPARE:          "帐号密码错误",
	ERROR_NOT_COMPARE_PASSWORD: "两次密码输入不一致",
	ERROR_FAIL_ENCRYPTION:      "加密失败",
	ERROR_NOT_EXIST_PRODUCT:    "该商品不存在",
	ERROR_NOT_EXIST_ADDRESS:    "该收货地址不存在",
	ERROR_EXIST_FAVORITE:       "已收藏该商品",

	ERROR_AUTH_CHECK_TOKEN_FAIL:       "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:    "Token已超时",
	ERROR_AUTH_TOKEN:                  "Token生成失败",
	ERROR_AUTH:                        "Token错误",
	ERROR_AUTH_INSUFFICIENT_AUTHORITY: "权限不足",

	ERROR_DATABASE: "数据库操作出错，请重试",

	ERROR_OSS: "OSS配置错误",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
