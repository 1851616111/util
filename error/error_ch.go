package error

import "errors"

var (
	ParseHttpRequestErr = errors.New("解析http请求失败")
	ValidateRequestErr  = errors.New("请求参数验证失败")
)
