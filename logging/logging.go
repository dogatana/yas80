package logging

import (
	"fmt"
	"slices"
	"yas80/filecontent"
	"yas80/internal/util"
)

type LogMessage interface {
	Message() string
	Error() string
}

type MessageType int

const (
	Err = iota
	Warn
	Info
)

var msgTypeName = map[MessageType]string{
	Err:  "[ERR]",
	Warn: "[WARN]",
	Info: "[INFO]",
}

type Message struct {
	Type    MessageType
	message string
	Context *filecontent.Context
}

func (m *Message) Message() string { return msgTypeName[m.Type] + m.message }
func (m *Message) Error() string {
	tn := msgTypeName[m.Type]
	if m.Context == nil {
		return fmt.Sprintf("%s %s", tn, m.message)
	} else {
		return fmt.Sprintf("%q:%d %s %s", m.Context.FileContent.Filename, m.Context.Line, tn, m.message)
	}
}

// 2 つのメッセージが等しいかどうか（メソッド）
func (m *Message) Equal(o *Message) bool {
	return Equal(m, o)
}

// 2 つのメッセージが等しいかどうか（関数）
func Equal(a, b *Message) bool {
	return a.Type == b.Type && a.message == b.message && a.Context.Equal(b.Context)
}

type ErrorMessage struct {
	message string
	Context *filecontent.Context
}

func (m *ErrorMessage) Message() string { return m.message }
func (m *ErrorMessage) Error() string {
	if m.Context == nil {
		return fmt.Sprintf("[ERROR] %s", m.message)
	} else {
		return fmt.Sprintf("%q:%d [ERROR] %s", m.Context.FileContent.Filename, m.Context.Line, m.message)
	}
}
func (m *ErrorMessage) Equal(o *ErrorMessage) bool {
	return m.message == o.message && m.Context.Equal(o.Context)
}

type WarningMessage struct {
	message string
	Context *filecontent.Context
}

func (m *WarningMessage) Message() string { return m.message }
func (m *WarningMessage) Error() string {
	if m.Context == nil {
		return fmt.Sprintf("[WARN] %s", m.message)
	} else {
		return fmt.Sprintf("%q:%d [WARN] %s", m.Context.FileContent.Filename, m.Context.Line, m.message)
	}
}
func (m *WarningMessage) Equal(o *WarningMessage) bool {
	return m.message == o.message && m.Context.Equal(o.Context)
}

type InfoMessage struct {
	message string
	Context *filecontent.Context
}

func (m *InfoMessage) Message() string { return m.message }
func (m *InfoMessage) Error() string {
	if m.Context == nil {
		return fmt.Sprintf("[INFO] %s", m.message)
	} else {
		return fmt.Sprintf("%q:%d [INFO] %s", m.Context.FileContent.Filename, m.Context.Line, m.message)
	}
}
func (m *InfoMessage) Equal(o *InfoMessage) bool {
	return m.message == o.message && m.Context.Equal(o.Context)
}

type Logger struct {
	messages   []*Message
	Errors     []LogMessage
	Warnings   []LogMessage
	Infomation []LogMessage
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Count() (int, int, int) {
	var e, w, i int
	for _, m := range l.messages {
		switch m.Type {
		case Err:
			e++
		case Warn:
			w++
		case Info:
			i++
		}
	}
	return e, w, i
}

func (l *Logger) ErrorCount() int {
	e, _, _ := l.Count()
	return e
}

func (l *Logger) Error(msg string, ctx *filecontent.Context) error {
	err := &ErrorMessage{message: msg, Context: ctx}
	l.Errors = append(l.Errors, err)
	m := &Message{Type: Err, message: msg, Context: ctx}
	l.messages = append(l.messages, m)
	return err
}

func (l *Logger) Warning(msg string, ctx *filecontent.Context) error {
	err := &WarningMessage{message: msg, Context: ctx}
	l.Warnings = append(l.Warnings, err)
	m := &Message{Type: Warn, message: msg, Context: ctx}
	l.messages = append(l.messages, m)
	return err
}
func (l *Logger) Info(msg string, ctx *filecontent.Context) error {
	err := &InfoMessage{message: msg, Context: ctx}
	l.Infomation = append(l.Infomation, err)
	m := &Message{Type: Info, message: msg, Context: ctx}
	l.messages = append(l.messages, m)
	return err
}

func messageCompare(a, b *Message) int {
	return 0
}

func (l *Logger) Sort() {
	slices.SortFunc(l.messages, messageCompare)
}

func (l *Logger) Uniq() {

}

// logMessage の表示
func (l *Logger) Print() {
	// Warning, Information の重複を削除
	if len(l.Errors) != 0 {
		fmt.Printf("%d errros\n", len(l.Errors))
		for _, e := range l.Errors {
			fmt.Println(e.Error())
		}
	}
	if len(l.Warnings) != 0 {
		fmt.Printf("%d warnings\n", len(l.Warnings))
		for _, e := range l.Warnings {
			fmt.Println(e.Error())
		}
	}
	if len(l.Infomation) != 0 {
		fmt.Printf("%d info\n", len(l.Infomation))
		for _, e := range l.Infomation {
			fmt.Println(e.Error())
		}
	}
}

// Warnin, Information から重複したメッセージを削除する
func (l *Logger) RemoveDupe() {
	l.Errors = l.removeDupedMessage(l.Errors)
	l.Warnings = l.removeDupedMessage(l.Warnings)
	l.Infomation = l.removeDupedMessage(l.Infomation)
	l.removeDupedMessageNew()
}

// 重複メッセージの削除（Logger.message 版)
func (l *Logger) removeDupedMessageNew() {
	if len(l.messages) < 2 {
		return
	}
	for i := 0; i < len(l.messages)-1; i++ {
		m := l.messages[i]
		if m == nil {
			continue
		}
		for j := i + 1; j < len(l.messages); j++ {
			if l.messages[j] == nil {
				continue
			}
			if m.Equal(l.messages[j]) {
				l.messages[j] = nil
			}
		}
	}
	l.messages = util.Filter(l.messages, func(m *Message) bool { return m != nil })
}

// removeDupe のサポート関数
func (l *Logger) removeDupedMessage(msgs []LogMessage) []LogMessage {
	for i := 0; i < len(msgs)-1; i++ {
		m := msgs[i]
		if m == nil {
			continue
		}
		for j := i + 1; j < len(msgs); j++ {
			if msgs[j] == nil {
				continue
			}
			switch m := m.(type) {
			case *ErrorMessage:
				if m.Equal(msgs[j].(*ErrorMessage)) {
					msgs[j] = nil
				}
			case *WarningMessage:
				if m.Equal(msgs[j].(*WarningMessage)) {
					msgs[j] = nil
				}
			case *InfoMessage:
				if m.Equal(msgs[j].(*InfoMessage)) {
					msgs[j] = nil
				}
			}
		}
	}
	result := []LogMessage{}
	for _, m := range msgs {
		if m != nil {
			result = append(result, m)
		}
	}
	return result
}
