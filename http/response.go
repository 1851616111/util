package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Responser is a util interface for prasing data like {"code": "message", data:""}.
type Responser interface {
	Error() error
	Code() int
	Message() string
	Data() interface{}
}

func ReadToTarget(src io.Reader, dst interface{}, errH func(Responser) error) error {
	var handler func(Responser) error = errH
	if handler == nil {
		handler = DefaultErrorHandler
	}

	rsp := NewResponser(dst, handler)
	if err := json.NewDecoder(src).Decode(rsp.(*container)); err != nil {
		return err
	}

	return rsp.Error()
}

// dst must be ptr type
func NewResponser(dst interface{}, errFn func(Responser) error) Responser {
	return &container{
		DataStruct:   dst,
		ErrorHandler: errFn,
	}
}

func (c *container) Code() int {
	return c.C
}

func (c *container) Message() string {
	return c.Msg
}

func (c *container) Data() interface{} {
	return c.DataStruct
}

func (c *container) Error() error {
	return c.ErrorHandler((Responser)(c))
}

type container struct {
	C          int         `json:"code"`
	Msg        string      `json:"message"`
	CDesc      string      `json:"codeDesc"`
	TotalCount int         `json:"totalCount"`
	DataStruct interface{} `json:"data"`

	ErrorHandler func(Responser) error
}

func DefaultErrorHandler(rsp Responser) error {
	if rsp.Code() == 0 || rsp.Code() == 200 {
		return nil
	}

	return fmt.Errorf("%s", rsp.Message())
}

func Response(w http.ResponseWriter, header int, body interface{}) {
	w.WriteHeader(header)
	fmt.Fprintf(w, "%v", body)
	return
}

func ResponseJson(w http.ResponseWriter, header int, body interface{}) error {
	w.WriteHeader(header)
	return json.NewEncoder(w).Encode(body)
}
