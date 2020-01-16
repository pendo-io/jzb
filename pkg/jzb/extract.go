package jzb

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

type JzbBody struct {
	Input []byte
}

// ExtractJson will attempt to extract the json body from the given JZB input
// if an error is encountered, io.Reader will be nil
func (j JzbBody) ExtractJson() (io.Reader, error) {
	var buf []byte
	var err error
	if buf, err = base64.RawURLEncoding.DecodeString(string(j.Input)); err != nil {
		return nil, errors.New(fmt.Sprintf("error decoding jzb: %s", err.Error()))
	}
	reader, err := zlib.NewReader(bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}
	jsonBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("corrupt jzb: %s", err.Error()))
	}
	jsonBuf := bytes.NewBuffer(jsonBytes)
	return jsonBuf, nil
}
