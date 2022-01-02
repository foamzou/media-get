package utils

import (
	"path"
	"strings"
)

func GetExtFromUrl(url string) string {
	ext := strings.Split(path.Ext(url), "?")[0]
	ext = strings.Split(ext, "#")[0]
	return strings.Split(ext, ".")[1]
}
