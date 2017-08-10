package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Responser is a util interface for prasing D like {"code": "message", D:""}.
type Responser interface {
	Error() error
	Code() int
	Message() string
	Data() []byte
}

func ReadToTarget(src io.Reader, dst interface{}, errH func(Responser) error) error {
	var handler func(Responser) error = errH
	if handler == nil {
		handler = DefaultErrorHandler
	}

	rsp := NewResponser(handler)
	if err := json.NewDecoder(src).Decode(rsp.(*container)); err != nil {
		return err
	}

	if err := json.Unmarshal(rsp.Data(), &(rsp.(*container).baseFiled)); err != nil {
		return err
	}

	if err := json.Unmarshal(rsp.Data(), dst); err != nil {
		return err
	}

	return rsp.Error()
}

// dst must be ptr type
func NewResponser(errFn func(Responser) error) Responser {
	return &container{
		errorHandler: errFn,
		RawMessage:   &json.RawMessage{},
	}
}

func (c *container) Code() int {
	return c.baseFiled.Code
}

func (c *container) Message() string {
	return c.baseFiled.Message
}

func (c *container) Data() []byte {
	if c.RawMessage == nil {
		return nil
	}

	return ([]byte)(*c.RawMessage)
}

func (c *container) Error() error {
	return c.errorHandler((Responser)(c))
}

type container struct {
	*json.RawMessage
	baseFiled struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	errorHandler func(Responser) error
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
