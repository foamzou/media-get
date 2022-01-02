package utils

import (
	"errors"
	"regexp"
	"strconv"
)

func RegexSingleMatch(input string, expression string) (string, error) {
	r, err := regexp.Compile(expression)
	if err != nil {
		return "", err
	}
	match := r.FindStringSubmatch(input)
	if len(match) != 2 {
		return "", errors.New("should match one only")
	}
	return match[1], nil
}

func RegexSingleMatchIgnoreError(input string, expression string, fallback string) string {
	str, err := RegexSingleMatch(input, expression)
	if err != nil {
		return fallback
	}
	return str
}

func RegexSingleMatchIntIgnoreError(input string, expression string, fallback int) int {
	str, err := RegexSingleMatch(input, expression)
	if err != nil {
		return fallback
	}
	n, err := strconv.Atoi(str)
	if err != nil {
		return fallback
	}
	return n
}
