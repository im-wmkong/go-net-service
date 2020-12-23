package types

import (
	"go-net-service/pkg/logger"
	"strconv"
	"time"
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
		logger.Error(err)
	}
	return i
}

func DateFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func DateParse(str string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		logger.Error(err)
	}
	return t
}
