package parser

import (
	__yyfmt__ "fmt"
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

func opString(e Node) string {
	indirect, ok := e.(*IndirectExpression)
	if ok {
		return indirect.String()
	}
	s := e.String()
	return trimParen(s)
}
func trimParen(s string) string {
	__yyfmt__.Printf("trim %q", s)
	if s[0] == '(' {
		s = s[1 : len(s)-1]
	}
	__yyfmt__.Printf("to %q\n", s)
	return s
}

func SetYYDebug(v int) {
	yyDebug = v
}
