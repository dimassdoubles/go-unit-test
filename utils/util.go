package utils

import (
	"strings"
)

func UcFirst(str string) string {
	if len(str) == 0 {
		return str
	}

	if len(str) == 1 {
		return strings.ToUpper(str[0:1])
	}

	return strings.ToUpper(str[0:1]) + str[1:]
}
