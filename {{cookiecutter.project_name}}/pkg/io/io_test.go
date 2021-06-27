package io

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteToFile(t *testing.T) {
	tempFile, _ := ioutil.TempFile("", "WriteToFile.txt")

	err := WriteToFile(tempFile.Name(), []byte("test string"))
	if assert.Nil(t, err) {
		assert.FileExists(t, tempFile.Name())
	}
	defer os.Remove(tempFile.Name())
}

func TestYamlToFile(t *testing.T) {
	tempFile, _ := ioutil.TempFile("", "DataToYamlFile.yaml")

	err := DataToYamlFile(tempFile.Name(), []byte("test string"))
	if assert.Nil(t, err) {
		assert.FileExists(t, tempFile.Name())
	}
	defer os.Remove(tempFile.Name())
}
