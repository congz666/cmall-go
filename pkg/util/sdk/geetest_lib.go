package sdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"errors"
	"time"
	"encoding/json"
	"math/rand"
	"crypto/md5"
	"crypto/hmac"
	"crypto/sha256"
)

/**
 * sdk lib包，核心逻辑。
 *
 * @author liuquan@geetest.com
 */
const (
	IS_DEBUG                          bool   = true // 调试开关，是否输出调试日志
	API_URL                           string = "http://api.geetest.com"
	REGISTER_URL                      string = "/register.php"
	VALIDATE_URL                      string = "/validate.php"
	JSON_FORMAT                       string = "1"
	NEW_CAPTCHA                       bool   = true
	HTTP_TIMEOUT_DEFAULT              int    = 5 // 单位：秒
	VERSION                           string = "golang-gin:3.1.0"
	GEETEST_CHALLENGE                 string = "geetest_challenge" // 极验二次验证表单传参字段 chllenge
	GEETEST_VALIDATE                  string = "geetest_validate"  // 极验二次验证表单传参字段 validate
	GEETEST_SECCODE                   string = "geetest_seccode"   // 极验二次验证表单传参字段 seccode
	GEETEST_SERVER_STATUS_SESSION_KEY string = "gt_server_status"  // 极验验证API服务状态Session Key
)

type GeetestLib struct {
	geetest_id  string // 公钥
	geetest_key string // 私钥
	libResult   *GeetestLibResult
}

func NewGeetestLib(geetest_id string, geetest_key string) *GeetestLib {
	return &GeetestLib{geetest_id, geetest_key, NewGeetestLibResult()}
}

func (g *GeetestLib) gtlog(msg string) {
	if IS_DEBUG {
		fmt.Println("gtlog: " + msg)
	}
}

/**
 * 验证初始化
 */
func (g *GeetestLib) Register(digestmod string, params map[string]string) *GeetestLibResult {
	g.gtlog(fmt.Sprintf("Register(): 开始验证初始化, digestmod=%s.", digestmod))
	origin_challenge := g.requestRegister(params)
	g.buildRegisterResult(origin_challenge, digestmod)
	g.gtlog(fmt.Sprintf("Register(): 验证初始化, lib包返回信息=%s.", g.libResult))
	return g.libResult
}

/**
 * 向极验发送验证初始化的请求，GET方式
 */
func (g *GeetestLib) requestRegister(params map[string]string) string {
	params["gt"] = g.geetest_id
	params["json_format"] = JSON_FORMAT
	params["sdk"] = VERSION
	register_url := API_URL + REGISTER_URL
	g.gtlog(fmt.Sprintf("requestRegister(): 验证初始化, 向极验发送请求, url=%s, params=%s.", register_url, params))
	resBody, err := g.httpGet(register_url, params)
	if err != nil {
		g.gtlog(fmt.Sprintf("requestRegister(): 验证初始化, 请求异常，后续流程走宕机模式, %s", err))
		return ""
	}
	g.gtlog(fmt.Sprintf("requestRegister(): 验证初始化, 与极验网络交互正常, 返回body=%s.", resBody))
	resMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(resBody), &resMap)
	if err != nil {
		g.gtlog(fmt.Sprintf("requestRegister(): 验证初始化, 解析json异常，后续流程走宕机模式, %s", err))
		return ""
	}
	return resMap["challenge"].(string)
}

/**
 * 构建验证初始化返回数据
 */
func (g *GeetestLib) buildRegisterResult(origin_challenge string, digestmod string) {
	// origin_challenge为空或者值为0代表失败
	if origin_challenge == "" || origin_challenge == "0" {
		// 本地随机生成32位字符串
		characters := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
		challenge := []byte{}
		for i := 0; i < 32; i++ {
			challenge = append(challenge, characters[rand.Intn(len(characters))])
		}
		dataMap := map[string]interface{}{
			"success":     0,
			"gt":          g.geetest_id,
			"challenge":   string(challenge),
			"new_captcha": NEW_CAPTCHA,
		}
		dataJson, _ := json.Marshal(dataMap)
		g.libResult.setAll(0, string(dataJson), "请求极验register接口失败，后续流程走宕机模式")
	} else {
		challenge := ""
		if digestmod == "md5" {
			challenge = g.md5_encode(origin_challenge + g.geetest_key)
		} else if digestmod == "sha256" {
			challenge = g.sha256_encode(origin_challenge + g.geetest_key)
		} else if digestmod == "hmac-sha256" {
			challenge = g.hmac_sha256_encode(origin_challenge, g.geetest_key)
		} else {
			challenge = g.md5_encode(origin_challenge + g.geetest_key)
		}
		dataMap := map[string]interface{}{
			"success":     1,
			"gt":          g.geetest_id,
			"challenge":   challenge,
			"new_captcha": NEW_CAPTCHA,
		}
		dataJson, _ := json.Marshal(dataMap)
		g.libResult.setAll(1, string(dataJson), "")
	}
}

/**
 * 正常流程下（即验证初始化成功），二次验证
 */
func (g *GeetestLib) SuccessValidate(challenge string, validate string, seccode string, params map[string]string) *GeetestLibResult {
	g.gtlog(fmt.Sprintf("SuccessValidate(): 开始二次验证 正常模式, challenge=%s, validate=%s, seccode=%s.", challenge, validate, seccode))
	if !g.checkParam(challenge, validate, seccode) {
		g.libResult.setAll(0, "", "正常模式，本地校验，参数challenge、validate、seccode不可为空")
	} else {
		response_seccode := g.requestValidate(challenge, validate, seccode, params)
		if response_seccode == "" {
			g.libResult.setAll(0, "", "请求极验validate接口失败")
		} else if response_seccode == "false" {
			g.libResult.setAll(0, "", "极验二次验证不通过")
		} else {
			g.libResult.setAll(1, "", "")
		}
	}
	g.gtlog(fmt.Sprintf("SuccessValidate(): 二次验证 正常模式, lib包返回信息=%s.", g.libResult))
	return g.libResult
}

/**
 * 异常流程下（即验证初始化失败，宕机模式），二次验证
 * 注意：由于是宕机模式，初衷是保证验证业务不会中断正常业务，所以此处只作简单的参数校验，可自行设计逻辑。
 */
func (g *GeetestLib) FailValidate(challenge string, validate string, seccode string) *GeetestLibResult {
	g.gtlog(fmt.Sprintf("FailValidate(): 开始二次验证 宕机模式, challenge=%s, validate=%s, seccode=%s.", challenge, validate, seccode))
	if !g.checkParam(challenge, validate, seccode) {
		g.libResult.setAll(0, "", "宕机模式，本地校验，参数challenge、validate、seccode不可为空.")
	} else {
		g.libResult.setAll(1, "", "")
	}
	g.gtlog(fmt.Sprintf("FailValidate(): 二次验证 宕机模式, lib包返回信息=%s.", g.libResult))
	return g.libResult
}

/**
 * 向极验发送二次验证的请求，POST方式
 */
func (g *GeetestLib) requestValidate(challenge string, validate string, seccode string, params map[string]string) string {
	params["seccode"] = seccode
	params["json_format"] = JSON_FORMAT
	params["challenge"] = challenge
	params["sdk"] = VERSION
	params["captchaid"] = g.geetest_id
	validate_url := API_URL + VALIDATE_URL
	g.gtlog(fmt.Sprintf("requestValidate(): 二次验证 正常模式, 向极验发送请求, url=%s, params=%s.", validate_url, params))
	resBody, err := g.httpPost(validate_url, params)
	if err != nil {
		g.gtlog(fmt.Sprintf("requestValidate(): 二次验证 正常模式, 请求异常, %s", err))
		return ""
	}
	g.gtlog(fmt.Sprintf("requestValidate(): 二次验证 正常模式, 与极验网络交互正常, 返回body=%s.", resBody))
	resMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(resBody), &resMap)
	if err != nil {
		g.gtlog(fmt.Sprintf("requestValidate(): 二次验证 正常模式, 解析json异常, %s", err))
		return ""
	}
	return resMap["seccode"].(string)
}

/**
 * 校验二次验证的三个参数，校验通过返回true，校验失败返回false
 */
func (g *GeetestLib) checkParam(challenge string, validate string, seccode string) bool {
	return !(challenge == "" || strings.TrimSpace(challenge) == "" || validate == "" || strings.TrimSpace(validate) == "" || seccode == "" || strings.TrimSpace(seccode) == "")
}

/**
 * 发送GET请求，获取服务器返回结果
 */
func (g *GeetestLib) httpGet(getUrl string, params map[string]string) (string, error) {
	q := url.Values{}
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
	}
	req, err := http.NewRequest(http.MethodGet, getUrl, nil)
	if err != nil {
		return "", errors.New("NewRequest fail")
	}
	req.URL.RawQuery = q.Encode()
	client := &http.Client{Timeout: time.Duration(HTTP_TIMEOUT_DEFAULT) * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode == 200 {
		return string(body), nil
	}
	return "", nil
}

/**
 * 发送POST请求，获取服务器返回结果
 */
func (g *GeetestLib) httpPost(postUrl string, params map[string]string) (string, error) {
	q := url.Values{}
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
	}
	req, err := http.NewRequest(http.MethodPost, postUrl, strings.NewReader(q.Encode()))
	if err != nil {
		return "", errors.New("NewRequest fail")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{Timeout: time.Duration(HTTP_TIMEOUT_DEFAULT) * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode == 200 {
		return string(body), nil
	}
	return "", nil
}

/**
 * md5 加密
 */
func (g *GeetestLib) md5_encode(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%x", h.Sum(nil))
}

/**
 * sha256加密
 */
func (g *GeetestLib) sha256_encode(value string) string {
	h := sha256.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%x", h.Sum(nil))
}

/**
 * hmac-sha256 加密
 */
func (g *GeetestLib) hmac_sha256_encode(value string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(value))
	return fmt.Sprintf("%x", h.Sum(nil))
}
