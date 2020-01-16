package jzb

import (
	"bytes"
	"compress/gzip"
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
	gzipWriter := gzip.NewWriter(&buf)
	if _, err := gzipWriter.Write(j.Input); err != nil {
		return nil, err
	}
	if err := gzipWriter.Flush(); err != nil {
		return nil, err
	}
	_ = gzipWriter.Close()
	var b64bytes []byte
	base64.StdEncoding.Encode(b64bytes, buf.Bytes())
	returnBuf := bytes.NewBuffer(b64bytes)
	return returnBuf, nil
}
