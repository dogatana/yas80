package logging

import (
	"fmt"
	"yas80/fileblock"
)

type LogMessage interface {
	Message() string
	Error() string
}
type ErrorMessage struct {
	message string
	Context *fileblock.Context
}

func (m *ErrorMessage) Message() string { return m.message }
func (m *ErrorMessage) Error() string {
	if m.Context == nil {
		return fmt.Sprintf("%q:%d [ERROR] %s", "???", -1, m.message)
	} else {
		return fmt.Sprintf("%q:%d [ERROR] %s", m.Context.FileBlock.Filename, m.Context.Line, m.message)
	}
}

type WarningMessage struct {
	message string
	Context *fileblock.Context
}

func (m *WarningMessage) Message() string { return m.message }
func (m *WarningMessage) Error() string {
	if m.Context == nil {
		return fmt.Sprintf("%q:%d [ERROR] %s", "???", -1, m.message)
	} else {
		return fmt.Sprintf("%q:%d [WARN] %s", m.Context.FileBlock.Filename, m.Context.Line, m.message)
	}
}

type InfoMessage struct {
	message string
	Context *fileblock.Context
}

func (m *InfoMessage) Message() string { return m.message }
func (m *InfoMessage) Error() string {
	if m.Context == nil {
		return fmt.Sprintf("%q:%d [ERROR] %s", "???", -1, m.message)
	} else {
		return fmt.Sprintf("%q:%d [INFO] %s", m.Context.FileBlock.Filename, m.Context.Line, m.message)
	}
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

func (l *Logger) Error(msg string, ctx *fileblock.Context) error {
	err := &ErrorMessage{message: msg, Context: ctx}
	l.Errors = append(l.Errors, err)
	return err
}

func (l *Logger) Warning(msg string, ctx *fileblock.Context) error {
	err := &WarningMessage{message: msg, Context: ctx}
	l.Warnings = append(l.Warnings, err)
	return err
}
func (l *Logger) Info(msg string, ctx *fileblock.Context) error {
	err := &InfoMessage{message: msg, Context: ctx}
	l.Infomation = append(l.Infomation, err)
	return err
}

func (l *Logger) Count() (int, int, int) {
	return len(l.Errors), len(l.Warnings), len(l.Infomation)
}

func (l *Logger) Print() {
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
