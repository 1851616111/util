package message

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sync"

	"fmt"
	httputil "github.com/1851616111/util/http"
)

var bytesPool *sync.Pool

func init() {
	bytesPool = &sync.Pool{}
	bytesPool.New = func() interface{} {
		return &bytes.Buffer{}
	}
}

var codeMsgM map[int]interface{} = map[int]interface{}{
	1000: "success",
	1001: ERR_SERVER_INNER_ERROR,
	1004: ERR_REQ_NOT_FOUND_ERROR,
}

//向用户返回错误信息
var ERR_SERVER_INNER_ERROR error = errors.New("Internal Server Error")
var ERR_REQ_NOT_FOUND_ERROR error = errors.New("Request Not Found")

const _Inner_Error = `{"code":1001,"message":"Internal Server Error"}`
const _Req_Not_Find = `{"code":1004,"message":"Request Not Found"}`
const _Param_Not_Find = `{"code":1004,"message":"Param %s Not Found"}`
const _Unauthorized = `{"code":1005,"message":"Unauthorized"}`
const SuccessCode = 1000

type Message struct {
	Code int         `json:"code"`
	Data interface{} `json:"data, omitempty"`
	Msg  string      `json:"message, omitempty""`
}

func Success(w http.ResponseWriter) {
	httputil.Response(w, 200, `{"code":1000,"message":"success"}`)
}

func SuccessI(w http.ResponseWriter, obj interface{}) error {
	return json.NewEncoder(w).Encode(obj)
}

func SuccessS(w http.ResponseWriter, s string) {
	httputil.Response(w, 200, fmt.Sprintf(`{"code":1000,"data":"%s",message":"success"}`, s))
}

func SuccessObj(w http.ResponseWriter, obj interface{}) error {
	return json.NewEncoder(w).Encode(Message{Code: 1000, Data: obj, Msg: "success"})
}

func InnerError(w http.ResponseWriter) {
	httputil.Response(w, 400, message(_Inner_Error))
}

func Unauthorized(w http.ResponseWriter) {
	httputil.Response(w, http.StatusUnauthorized, message(_Unauthorized))
}

func NotFoundError(w http.ResponseWriter) {
	httputil.Response(w, 404, message(_Req_Not_Find))
}

//func ParamNotFound(w http.ResponseWriter, param string) {
//	httputil.Response(w, 400, message(fmt.Errorf(_Param_Not_Find, param)))
//}

func Render(w http.ResponseWriter, msg interface{}) {
	httputil.ResponseJson(w, 200, Message{Code: 1000, Data: msg, Msg: "success"})
}

func Error(w http.ResponseWriter, err error) {
	httputil.ResponseJson(w, 400, Message{Code: 1004, Msg: err.Error()})
}

//{"code":1000,"data":null,"message":"success"}
func generateData() *bytes.Buffer {
	buf := message(`{"code":1000,"message":"success"}`)
	return buf
}

func ReadBody(body io.Reader) ([]byte, error) {
	tmp := bytesPool.Get().(*bytes.Buffer)
	if _, err := tmp.ReadFrom(body); err != nil {
		return nil, err
	}

	data := tmp.Bytes()
	if tmp.Len() > 0 && data[tmp.Len()-1] == 0xa {
		data = data[:tmp.Len()-1]
	}

	return data, nil
}

func message(msgs ...string) *bytes.Buffer {
	buf := bytesPool.Get().(*bytes.Buffer)
	if len(msgs) > 0 {
		for _, m := range msgs {
			buf.WriteString(m)
		}
	}

	return buf
}
