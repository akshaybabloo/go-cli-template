package io

import (
	"os"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

// WriteToFile Writes byte array to a file path
func WriteToFile(filePath string, data []byte) error {
	fs := afero.NewOsFs()
	appFs := afero.Afero{Fs: fs}

	err := appFs.WriteFile(filePath, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// DataToYamlFile creates a YAML file
func DataToYamlFile(fileName string, data interface{}) error {
	marshal, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	err = WriteToFile(fileName, marshal)
	if err != nil {
		return err
	}

	return nil
}
