package parser

import (
	"strconv"
	"strings"
)

func parseInt(s string) (int64, error) {
	str := strings.ReplaceAll(s, "_", "")
	length := len(str)
	switch {
	case length >= 3 && str[0] == '0':
		return strconv.ParseInt(str, 0, 0)
	case length >= 2 && str[0] == '$':
		return strconv.ParseInt("0x"+str[1:length], 0, 0)
	case length >= 2 && (str[length-1] == 'h' || str[length-1] == 'H'):
		return strconv.ParseInt(str[0:length-1], 16, 0)
	default:
		return strconv.ParseInt(str, 0, 0)
	}
}
