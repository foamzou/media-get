package utils

import (
	"math"
	"strconv"
	"strings"
)

func DurationStr2Second(durationStr string) int {
	l := strings.Split(durationStr, ":")
	second := 0
	lenOfL := len(l)
	for i := lenOfL; i > 0; i-- {
		base, err := strconv.Atoi(l[i-1])
		if err != nil {
			return 0
		}

		second += base * int(math.Pow(float64(60), float64(lenOfL-i)))
	}
	return second
}
