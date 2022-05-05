package utils

import (
	"strings"
)

func GetExtFromUrl(url string) string {
	ext := strings.Split(url, "?")[0]
	ext = strings.Split(ext, "#")[0]
	extGroup := strings.Split(ext, "/")

	if len(extGroup) <= 1 {
		return ""
	}

	extGroup = strings.Split(extGroup[len(extGroup)-1], ".")
	if len(extGroup) > 1 {
		return extGroup[len(extGroup)-1]
	}
	return ""
}
