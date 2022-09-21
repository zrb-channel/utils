package utils

import (
	"strconv"
	"time"
)

func Millisecond() string {
	return strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
}

func Timestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
