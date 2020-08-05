package sdk

import "fmt"

/**
 * sdk lib包的返回结果信息。
 *
 * @author liuquan@geetest.com
 */
type GeetestLibResult struct {
	Status int
	Data   string
	Msg    string
}

func NewGeetestLibResult() *GeetestLibResult {
	return &GeetestLibResult{0, "", ""}
}

func (g *GeetestLibResult) setAll(status int, data string, msg string) {
	g.Status = status
	g.Data = data
	g.Msg = msg
}

func (g *GeetestLibResult) String() string {
	return fmt.Sprintf("GeetestLibResult{Status=%s, Data=%s, Msg=%s}", g.Status, g.Data, g.Msg)
}
