package message

import "errors"

var (
	ErrParseHttpJsonReq error = errors.New("解析http请求异常")
)
