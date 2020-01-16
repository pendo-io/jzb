package cfg

import (
	"os"
	"testing"
)

func TestCommandLineArguments_Validate(t *testing.T) {
	type fields struct {
		InputPath  string
		OutputPath string
		Create     bool
		Extract    bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "No error when create option is selected",
			fields: fields{
				InputPath:  "-",
				OutputPath: "",
				Create:     true,
				Extract:    false,
			},
			wantErr: false,
		},
		{
			name: "No error when extract option is selected",
			fields: fields{
				InputPath:  "-",
				OutputPath: "",
				Create:     false,
				Extract:    true,
			},
			wantErr: false,
		},
		{
			name: "error when both create and extract options are selected",
			fields: fields{
				InputPath:  "-",
				OutputPath: "",
				Create:     true,
				Extract:    true,
			},
			wantErr: true,
		},
		{
			name: "error when neither create and extract options are selected",
			fields: fields{
				InputPath:  "-",
				OutputPath: "",
				Create:     false,
				Extract:    false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CommandLineArguments{
				InputPath:  tt.fields.InputPath,
				OutputFile: tt.fields.OutputPath,
				Create:     tt.fields.Create,
				Extract:    tt.fields.Extract,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommandLineArguments_Validate_File(t *testing.T) {
	if err := os.Mkdir("/tmp/testdir", os.ModeDir); err != nil {
		t.Fatal("could not set up test")
	}
	if _, err := os.Create("/tmp/test2.json"); err != nil {
		t.Fatal("could not set up test")
	}
	defer func() {
		_ = os.Remove("/tmp/testdir")
		_ = os.Remove("/tmp/test2.json")
	}()

	type fields struct {
		InputPath  string
		OutputPath string
		Create     bool
		Extract    bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "error when input path is a directory",
			fields: fields{
				InputPath:  "/tmp/testdir",
				OutputPath: "",
				Create:     true,
				Extract:    false,
			},
			wantErr: true,
		},
		{
			name: "input file does not exist",
			fields: fields{
				InputPath:  "/tmp/test.json",
				OutputPath: "",
				Create:     true,
				Extract:    false,
			},
			wantErr: true,
		},
		{
			name: "output file already exists",
			fields: fields{
				InputPath:  "_",
				OutputPath: "/tmp/test2.json",
				Create:     false,
				Extract:    true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CommandLineArguments{
				InputPath:  tt.fields.InputPath,
				OutputFile: tt.fields.OutputPath,
				Create:     tt.fields.Create,
				Extract:    tt.fields.Extract,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}
