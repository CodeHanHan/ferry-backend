package stringutil

import "strings"

func Join(arr ...string) string {
	buf := new(strings.Builder)
	for _, s := range arr {
		buf.WriteString(s)
	}
	return buf.String()
}
