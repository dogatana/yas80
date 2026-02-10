package filecontent

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
	"yas80/internal/util"
)

type FileContent struct {
	Filename string
	Content  []byte
	lines    []string
}

func (fc *FileContent) String() string {
	var content string
	if fc.Content == nil {
		content = "<nil>"
	} else {
		content = fmt.Sprintf("[]byte(len=%d)", len(fc.Content))
	}
	return fmt.Sprintf("FileBlock{FileName: %s, Content: %s}", fc.Filename, content)
}

// 指定行を取得
func (fc *FileContent) GetLine(line int) (string, error) {
	fc.setupLines()

	if line < 1 || line > len(fc.lines) {
		return "", fmt.Errorf("out of range %d / %d", line, len(fc.lines))
	}
	return fc.lines[line-1], nil
}

// Content []byte から lines []string を生成
func (fc *FileContent) setupLines() {
	if fc.lines == nil {
		fc.lines = strings.Split(string(fc.Content), "\n")
		// fc.Content = nil // ガベージコレクタ処理用に回す
	}
}

func (fc *FileContent) LineCount() int {
	fc.setupLines()
	return len(fc.lines)
}

func NewFromString(filename string, content string) (*FileContent, error) {
	return &FileContent{Filename: filename, Content: []byte(content)}, nil
}

func NewFromReader(filename string, reader io.Reader) (*FileContent, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return &FileContent{Filename: filename, Content: data}, nil

}
func NewFromFile(filename string) (*FileContent, error) {
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
		if data, err = util.ShiftJisToUtf8(data); err != nil {
			return nil, err
		}
	}

	// BOM があれば削除
	if bytes.HasPrefix(data, []byte{0xef, 0xbb, 0xbf}) {
		data = data[3:]
	}
	// CR/LF を LF へ
	data = bytes.ReplaceAll(data, []byte{13, 10}, []byte{10})
	if data[len(data)-1] != 10 {
		data = append(data, 10)
	}
	return &FileContent{Filename: filename, Content: data}, nil
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
