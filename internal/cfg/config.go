package cfg

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type CommandLineArguments struct {
	InputPath  string
	OutputFile string
	Create     bool
	Extract    bool
	Version    bool
}

func (c CommandLineArguments) Validate() error {
	if c.Version {
		return nil
	}
	builder := &strings.Builder{}
	if (c.Create && c.Extract) || (!c.Create && !c.Extract) {
		builder.WriteString("must choose one -c or -x\n")
	}
	if c.InputPath != "-" {
		if file, err := os.Stat(c.InputPath); os.IsNotExist(err) {
			builder.WriteString(fmt.Sprintf("file %s does not exist\n", c.InputPath))
		} else if os.IsPermission(err) {
			builder.WriteString(fmt.Sprintf("permission denied reading file %s\n", c.InputPath))
		} else {
			if file.IsDir() {
				builder.WriteString(fmt.Sprintf("%s is a directory. expecting a file\n", c.InputPath))
			}
		}
	}
	if len(c.OutputFile) > 0 {
		if file, err := os.Stat(c.OutputFile); err == nil {
			if file.IsDir() {
				builder.WriteString(fmt.Sprintf("%s is a directory. point to an output file\n", c.OutputFile))
			} else {
				builder.WriteString(fmt.Sprintf("%s already exists\n", c.OutputFile))
			}
		}
	}
	msg := builder.String()
	if len(msg) > 0 {
		return errors.New(msg)
	}
	return nil
}
