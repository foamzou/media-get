package utils

import (
	"strconv"
	"time"
)

func GenReqID() string {
	ms := time.Now().UnixNano() / 1e6
	return Md5(strconv.FormatInt(ms, 10))
}
