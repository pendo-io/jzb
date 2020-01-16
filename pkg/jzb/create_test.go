package jzb

import (
	"io"
	"io/ioutil"
	"testing"
)

func TestJsonBody_CreateJZB(t *testing.T) {
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
			name: "invalid json throws an error",
			fields: fields{
				Input: []byte(`{ "test": "test" }}`),
			},
			wantErr: true,
		},
		{
			name: "can create a valid jzb",
			fields: fields{
				Input: []byte(`{ "test": "test" }`),
			},
			wantErr: false,
			validate: func(read io.Reader) bool {
				bytes, err := ioutil.ReadAll(read)
				if err != nil {
					return false
				}
				output, err := JzbBody{Input: bytes}.ExtractJson()
				if err != nil {
					return false
				}
				outputJson, err := ioutil.ReadAll(output)
				if err != nil {
					return false
				}
				return string(outputJson) == `{ "test": "test" }`
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JsonBody{
				Input: tt.fields.Input,
			}
			read, err := j.CreateJZB()
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJZB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !tt.validate(read) {
				t.Errorf("failed to validate for case %s", tt.name)
			}
		})
	}
}
