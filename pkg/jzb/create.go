package jzb

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
)

type JsonBody struct {
	Input []byte
}

// CreateJZB will create a JZB from the given JSON body
// if an error is encountered, the reader returned will be nil
func (j JsonBody) CreateJZB() (io.Reader, error) {
	if !json.Valid(j.Input) {
		return nil, errors.New("invalid json body")
	}
	var buf bytes.Buffer
	zWriter := zlib.NewWriter(&buf)
	if _, err := zWriter.Write(j.Input); err != nil {
		return nil, err
	}
	if err := zWriter.Flush(); err != nil {
		return nil, err
	}
	_ = zWriter.Close()
	b64 := base64.RawURLEncoding.EncodeToString(buf.Bytes())
	returnBuf := bytes.NewBuffer([]byte(b64))
	return returnBuf, nil
}
