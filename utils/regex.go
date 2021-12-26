package utils

import (
	"errors"
	"regexp"
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
