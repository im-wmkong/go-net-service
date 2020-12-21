package types

import (
	"go-net-service/pkg/logger"
	"strconv"
)

func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}
