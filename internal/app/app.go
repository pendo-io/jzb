package app

import (
	"github.com/pendo-io/jzb/internal/cfg"
	"github.com/pendo-io/jzb/pkg/jzb"
	"io/ioutil"
	"os"
)

func Execute(config cfg.CommandLineArguments) error {
	var bytes []byte
	var err error
	if config.InputPath != "-" {
		bytes, err = ioutil.ReadFile(config.InputPath)
		if err != nil {
			return err
		}
	} else {
		bytes, err = ioutil.ReadAll(os.Stdin)
	}
	if config.Create {
		return create(config, bytes)
	} else if config.Extract {
		return extract(config, bytes)
	}
	return nil
}

func create(config cfg.CommandLineArguments, input []byte) error {
	jsonBody := jzb.JsonBody{
		Input: input,
	}
	read, err := jsonBody.CreateJZB()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(read)
	if err != nil {
		return err
	}
	return write(config, bytes)
}

func extract(config cfg.CommandLineArguments, input []byte) error {
	jzbBody := jzb.JzbBody{
		Input: input,
	}
	read, err := jzbBody.ExtractJson()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(read)
	if err != nil {
		return err
	}
	return write(config, bytes)
}

func write(config cfg.CommandLineArguments, bytes []byte) error {
	if len(config.OutputFile) > 0 {
		err := ioutil.WriteFile(config.OutputFile, bytes, 0644)
		if err != nil {
			return err
		}
	} else {
		println(string(bytes))
	}
	return nil
}
