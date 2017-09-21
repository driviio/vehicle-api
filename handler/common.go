package handler

import (
	"strconv"
)

func ParseInt64(val string) int64 {
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func ParseInt(val string) int {
	i, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return 0
	}
	return int(i)
}
