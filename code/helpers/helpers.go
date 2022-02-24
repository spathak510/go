package helpers

import (
	"strconv"
)

func Int64ToString(inputNum int64) string {
	return strconv.FormatInt(inputNum, 10)
}
