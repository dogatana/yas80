package filecontent

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
	"yas80/errcode"
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

// 総行数を返す
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

	data, err = formatBytesInput(data)
	if err != nil {
		return nil, err
	}
	return &FileContent{Filename: filename, Content: data}, nil
}

func NewFromFile(filename string) (*FileContent, error) {
	var data []byte
	var err error

	if data, err = os.ReadFile(filename); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf(errcode.EFILE_NOT_FOUND, filename)
		}
		return nil, fmt.Errorf(errcode.EFILE_ERR, filename, err.Error())
	}

	if data, err = formatBytesInput(data); err != nil {
		return nil, fmt.Errorf(errcode.EFILE_ERR, filename, err.Error())
	}

	return &FileContent{Filename: filename, Content: data}, nil
}

// []byte のエンコード、改行を整形する
func formatBytesInput(data []byte) ([]byte, error) {
	var err error

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

	// 最後に改行がなければ追加
	if len(data) == 0 {
		data = []byte{10}
	} else if data[len(data)-1] != 10 {
		data = append(data, 10)
	}

	return data, nil
}
