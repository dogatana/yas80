package logging

import (
	"fmt"
	"yas80/filecontent"
)

type LogMessage interface {
	Message() string
	Error() string
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
	Errors     []LogMessage
	Warnings   []LogMessage
	Infomation []LogMessage
	Filename   string
}

func New(filename string) *Logger {
	return &Logger{Filename: filename}
}

func (l *Logger) Error(msg string, ctx *filecontent.Context) error {
	err := &ErrorMessage{message: msg, Context: ctx}
	l.Errors = append(l.Errors, err)
	return err
}

func (l *Logger) Warning(msg string, ctx *filecontent.Context) error {
	err := &WarningMessage{message: msg, Context: ctx}
	l.Warnings = append(l.Warnings, err)
	return err
}
func (l *Logger) Info(msg string, ctx *filecontent.Context) error {
	err := &InfoMessage{message: msg, Context: ctx}
	l.Infomation = append(l.Infomation, err)
	return err
}

func (l *Logger) Count() (int, int, int) {
	return len(l.Errors), len(l.Warnings), len(l.Infomation)
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
	l.Warnings = l.removeDupedMessage(l.Warnings)
	l.Infomation = l.removeDupedMessage(l.Infomation)
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
