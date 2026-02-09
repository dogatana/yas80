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

// Map
func Map[T any](s []T, f func(T) T) []T {
	out := make([]T, len(s))
	for i, v := range s {
		out[i] = f(v)
	}
	return out
}

// Stack
type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}
func (s *Stack[T]) Len() int {
	return len(s.data)
}
func (s *Stack[T]) Empty() bool {
	return len(s.data) == 0
}
func (s *Stack[T]) Top() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
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
