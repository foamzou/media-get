package utils

import (
	"strconv"
	"strings"
)

func ConvertString2Int(numberStr string, fallback int) int {
	numberStr = strings.Split(numberStr, ".")[0]
	n, err := strconv.Atoi(numberStr)
	if err != nil {
		return fallback
	}
	return n
}
