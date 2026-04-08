package util

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/dogatana/yas80/errcode"

	"github.com/mattn/go-runewidth"
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
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	out := make([]T2, len(s))
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

// Orderd Map
type OrderedMap[K comparable, V any] struct {
	keys []K
	M    map[K]V
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{M: make(map[K]V)}
}

func (o *OrderedMap[K, V]) Set(k K, v V) {
	if _, exists := o.M[k]; !exists {
		o.keys = append(o.keys, k)
	}
	o.M[k] = v
}

func (o *OrderedMap[K, V]) Keys() []K { return o.keys }
func (o *OrderedMap[K, V]) Get(k K) (V, bool) {
	v, ok := o.M[k]
	return v, ok
}

// path 区切りを / にして返す
func SlashPath(path string) string {
	p := filepath.FromSlash(path)
	return filepath.ToSlash((p))
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
		return nil, errors.New(errcode.EFILE_ENCODING)
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

// TextObject を表示幅（34）に合わせて切り詰める
func TruncateWithEllipsis(s string, width int) string {
	rw := runewidth.StringWidth(s)
	if rw <= width {
		return s + strings.Repeat(" ", width-rw)
	}

	// 省略記号 "..." の幅は 3
	limit := width - 3
	if limit <= 0 {
		// 幅が小さすぎる場合は "..." のみ返す
		return "..."
	}

	out := ""
	w := 0

LOOP:
	for _, r := range s {
		rw := runewidth.RuneWidth(r)
		switch {
		case w+rw > limit:
			out += "... "
			break LOOP
		case w+rw == limit:
			out += string(r) + "..."
			break LOOP
		default:
			out += string(r)
			w += rw
		}
	}
	return out
}

// 文字列が ASCII 0x20-0x7e の範囲がどうかを返す
func IsAsciiString(name string) bool {
	for _, c := range name {
		if c < ' ' || c > 0x7e {
			return false
		}
	}
	return true
}

func PrintStructFields(v any) {
	// 構造体名
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	name := t.Name()

	// フィールド名、値
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}

	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		field := rt.Field(i) // フィールド情報（名前・タグなど）
		if field.Name[0] > 'Z' {
			continue // 非公開フィールドは無視
		}
		value := rv.Field(i) // 実際の値
		if s, ok := value.Interface().(string); ok {
			fmt.Printf("%s.%s = %q\n", name, field.Name, s)
		} else {
			fmt.Printf("%s.%s = %v\n", name, field.Name, value.Interface())
		}
	}
}

// 定義用匿名ラベル
func IsAnonDef(name string) bool {
	names := map[string]bool{
		"@@": true,
		"@0": true,
		"@1": true,
		"@2": true,
		"@3": true,
		"@4": true,
		"@5": true,
		"@6": true,
		"@7": true,
		"@8": true,
		"@9": true,
	}
	return names[name]
}

// 参照用匿名ラベル
func IsAnonRef(name string) bool {
	name = strings.ToUpper(name)
	if name[0] != '@' {
		return false
	}
	if len(name) == 2 && (name[1] == 'F' || name[1] == 'B') {
		return true
	}
	if len(name) == 3 && '0' <= name[1] && name[1] <= '9' && (name[2] == 'F' || name[2] == 'B') {
		return true
	}
	return false
}
