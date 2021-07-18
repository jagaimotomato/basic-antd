package response

import "time"

type Response struct {
	RequestId   string      `json:"request_id"`
	RequestTime string      `json:"request_time"`
	Code        int         `json:"code" example:"200"`
	Msg         string      `json:"msg,omitempty"`
	Result      interface{} `json:"result"`
	Timestamp   int64       `json:"timestamp"`
}

type Page struct {
	List      interface{} `json:"list"`
	Count     int64       `json:"count"`
	PageIndex int         `json:"page_index"`
	PageSize  int         `json:"page_size"`
}

type List struct {
	List  interface{} `json:"list"`
	Count int64       `json:"count"`
}

func (res *Response) Success() *Response {
	res.Code = 200
	res.Timestamp = time.Now().UnixNano() / 1e6
	res.RequestTime = time.Now().Format("2006-01-02 15:04:05")
	return res
}

func (res *Response) Error(code int) *Response {
	res.Code = code
	res.Timestamp = time.Now().UnixNano() / 1e6
	res.RequestTime = time.Now().Format("2006-01-02 15:04:05")
	return res
}

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	// CodeParamErr 各种奇怪参数错误
	CodeParamErr = 40001
)
