package utils

import (
	"os"
	"path"
	"strings"
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

func FilterUnexpectedChar(path string) string {
	invalidChars := []string{" ", ".", "/", "?", "\\", ":", "*", "\"", "<", ">", "|"}
	for _, char := range invalidChars {
		path = strings.ReplaceAll(path, char, "")
	}
	return path
}
