package file

import (
	"bytes"
	httput "github.com/1851616111/util/http"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func PostFile(url, field, fileName string, rc io.Reader) (*http.Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile(field, fileName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, rc)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	return http.Post(url, contentType, bodyBuf)
}

func GetHttpImage(targetFile string, spec *httput.HttpSpec) error {
	rsp, err := httput.Send(spec)
	if err != nil {
		return err
	}

	img, err := jpeg.Decode(rsp.Body)
	if err != nil {
		return err
	}

	target, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer target.Close()

	return jpeg.Encode(target, img, &jpeg.Options{jpeg.DefaultQuality})
}
