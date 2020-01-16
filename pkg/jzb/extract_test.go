package jzb

import (
	"io"
	"io/ioutil"
	"testing"
)

func TestJzbBody_ExtractJson(t *testing.T) {
	type fields struct {
		Input []byte
	}
	tests := []struct {
		name     string
		fields   fields
		wantErr  bool
		validate func(read io.Reader) bool
	}{
		{
			name: "invalid base64 results in error",
			fields: fields{
				Input: []byte(".3"),
			},
			wantErr: true,
		},
		{
			name: "corrupt jzb results in error",
			fields: fields{
				Input: []byte("bm90IGEgemlwCg"),
			},
			wantErr: true,
		},
		{
			name: "can extract json",
			fields: fields{
				Input: []byte(`eJyqVlAqSS0uUbKC0gq1XAAAAAD__wEAAP__OtwFpQ`),
			},
			wantErr: false,
			validate: func(read io.Reader) bool {
				bytes, err := ioutil.ReadAll(read)
				if err != nil {
					return false
				}
				output, err := JsonBody{Input: bytes}.CreateJZB()
				if err != nil {
					return false
				}
				outputJzb, err := ioutil.ReadAll(output)
				if err != nil {
					return false
				}
				return string(outputJzb) == `eJyqVlAqSS0uUbKC0gq1XAAAAAD__wEAAP__OtwFpQ`
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JzbBody{
				Input: tt.fields.Input,
			}
			read, err := j.ExtractJson()
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !tt.validate(read) {
				t.Errorf("failed to validate for case %s", tt.name)
			}
		})
	}
}
