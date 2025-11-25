package fileblock

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type FileBlock struct {
	Filename string
	Content  []byte
}

func (fb *FileBlock) String() string {
	var content string
	if fb.Content == nil {
		content = "<nil>"
	} else {
		content = fmt.Sprintf("[]byte(len=%d)", len(fb.Content))
	}
	return fmt.Sprintf("FileBlock{FileName: %s, Content: %s}", fb.Filename, content)
}

func New(filename string, content []byte) *FileBlock {
	return &FileBlock{Filename: filename, Content: content}
}

func NewFromReader(filename string, reader io.Reader) (*FileBlock, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return &FileBlock{Filename: filename, Content: data}, nil

}
func NewFromFile(filename string) (*FileBlock, error) {
	var data []byte
	var err error

	if err = checkFile(filename); err != nil {
		return nil, err
	}

	if data, err = os.ReadFile(filename); err != nil {
		return nil, err
	}

	if !utf8.Valid(data) {
		// cp932 と仮定し utf8 へ変換
		if data, err = shiftJisToUtf8(data); err != nil {
			return nil, err
		}
	}

	// BOM があれば削除
	if hasBOM(data) {
		data = data[3:]
	}
	// CR/LF を LF へ
	data = bytes.ReplaceAll(data, []byte{13, 10}, []byte{10})
	return &FileBlock{Filename: filename, Content: data}, nil
}

func checkFile(filename string) error {
	if st, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file not found: %s", filename)
		}
		return err
	} else if st.IsDir() {
		return fmt.Errorf("not a file: %s", filename)
	}
	return nil
}

func shiftJisToUtf8(input []byte) ([]byte, error) {
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

func hasBOM(data []byte) bool {
	if len(data) < 3 {
		return false
	}
	return data[0] == 0xef && data[1] == 0xbb && data[2] == 0xbf
}
