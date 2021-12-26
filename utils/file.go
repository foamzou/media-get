package utils

import (
	"io/ioutil"
	"os"
	"path"
)

func ModifyFileExt(filename, newExt string) string {
	ext := path.Ext(filename)
	return filename[0:len(filename)-len(ext)] + "." + newExt
}

func GetCurrentDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

func ReadFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
