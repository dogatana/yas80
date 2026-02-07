package util

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// Filter
func Filter[T any](s []T, f func(T) bool) []T {
	out := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

// Count
func Count[T any](s []T, f func(T) bool) int {
	out := 0
	for _, v := range s {
		if f(v) {
			out++
		}
	}
	return out
}

// utf-8 []byte を Shift-JIS []byte へ変換
func ShiftJisToUtf8(input []byte) ([]byte, error) {
	// Shift_JIS → UTF-8
	reader := transform.NewReader(bytes.NewReader(input), japanese.ShiftJIS.NewDecoder())
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	// u+fffd があれば正しく変換できていない
	if strings.ContainsRune(string(data), '\ufffd') {
		return nil, fmt.Errorf("unknown encoding")
	}
	return data, err
}

// utf-8 string を Shift-JIS []byte へ変換
func Utf8ToShiftJis(input string) ([]byte, error) {
	reader := transform.NewReader(strings.NewReader(input), japanese.ShiftJIS.NewEncoder())
	out, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return out, err
}
