package http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ContentType string

var ContentType_JSON ContentType = "application/json"
var ContentType_FORM ContentType = "application/x-www-form-urlencoded"

type HttpSpec struct {
	URL         string            `json:"url"`
	Method      string            `json:"method"`
	ContentType ContentType       `json:"content_type"`
	URLParams   *Params           `json:"url_params"`
	BodyParams  *Body             `json:"body_params"`
	BodyObject  interface{}       `json:"body_obj"`
	Header      map[string]string `json:"header"`
	BasicAuth   *BasicAuth        `json:"basicauth"`
}

func NewRequest(spec *HttpSpec) (*http.Request, error) {
	var req *http.Request
	var body io.Reader = nil
	var err error
	switch spec.Method {
	case "POST":
		if (spec.BodyObject != nil) || (spec.BodyParams != nil && len(*spec.BodyParams) > 0) {
			var target interface{}
			if spec.BodyObject != nil {
				target = spec.BodyObject
			} else {
				target = spec.BodyParams
			}

			switch spec.ContentType {
			case ContentType_JSON:
				buf := &bytes.Buffer{}
				encoder := json.NewEncoder(buf)
				encoder.SetEscapeHTML(false)

				if err := encoder.Encode(target); err != nil {
					return nil, err
				}
				body = io.Reader(buf)
			case ContentType_FORM:
				v := url.Values{}
				for key, value := range *spec.BodyParams {
					if s, ok := value.(string); ok {
						v.Add(key, s)
					}
				}
				body = ioutil.NopCloser(strings.NewReader(v.Encode()))
			}
		}
	case "GET":
	}

	urlStr := ([]byte)(spec.URL)
	if spec.URLParams != nil {
		urlStr = append(urlStr, "?"...)
		urlStr = append(urlStr, spec.URLParams.String()...)
	}

	if req, err = http.NewRequest(spec.Method, string(urlStr), body); err != nil {
		return nil, err
	}

	if len(spec.Header) > 0 {
		for k, v := range spec.Header {
			req.Header.Set(k, v)
		}
	}

	if spec.BasicAuth != nil {
		req.SetBasicAuth(spec.BasicAuth.User, spec.BasicAuth.Password)
	}

	req.Header.Set("Content-Type", string(spec.ContentType))

	return req, nil
}

func ReadJsonObj(spec *HttpSpec, obj interface{}) error {
	req, err := NewRequest(spec)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	return json.NewDecoder(req.Body).Decode(obj)
}

type BasicAuth struct {
	User     string
	Password string
}

type Body map[string]interface{}

func NewBody() *Body {
	return &Body{}
}

func (p *Body) Add(key string, value interface{}) *Body {
	map[string]interface{}(*p)[key] = value
	return p
}
