package stringutil

import (
	"strconv"
	"strings"
)

func Int2String(i int) string {
	return strconv.Itoa(i)
}

func Join(arr ...string) string {
	buf := new(strings.Builder)
	for _, s := range arr {
		buf.WriteString(s)
	}
	return buf.String()
}
