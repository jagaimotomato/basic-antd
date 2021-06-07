package app

import "time"

type Response struct {
	RequestId   string      `json:"requestId"`
	RequestTime string      `json:"requestTime"`
	Code        int         `json:"code" example:"200"`
	Message     string      `json:"message,omitempty"`
	Result      interface{} `json:"result"`
	Timestamp   int64       `json:"timestamp"`
}

type Page struct {
	List      interface{} `json:"list"`
	Count     int64       `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
}

type List struct {
	List  interface{} `json:"list"`
	Count int64       `json:"count"`
}

func (res *Response) ReturnSuccess() *Response {
	res.Code = 200
	res.Timestamp = time.Now().UnixNano() / 1e6
	res.RequestTime = time.Now().Format("2006-01-02 15:04:05")
	return res
}

func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	res.Timestamp = time.Now().UnixNano() / 1e6
	res.RequestTime = time.Now().Format("2006-01-02 15:04:05")
	return res
}
